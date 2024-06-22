package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/progate-hackathon-ari/backend/cmd/config"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/repository"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/s3"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type GameInteractor struct {
	repo    repository.DataAccess
	s3      s3.S3
	bedrock bedrock.Bedrock
	client  *Client
}

func NewGameInteractor(ws *websocket.Conn, repo repository.DataAccess, s3 s3.S3, bedrock bedrock.Bedrock) *GameInteractor {
	return &GameInteractor{
		repo:    repo,
		s3:      s3,
		bedrock: bedrock,
		client: &Client{
			ws: ws,
		},
	}
}

func (i *GameInteractor) CountDown(ctx context.Context, count int) error {
	return Counter(i.client.info.RoomID, count)
}

type Dummy struct {
	State NextState `json:"state"`
}

type NextState string

const (
	StateStartGame NextState = "start_game"
	StateNextRound NextState = "next_round"
	StateGameEnd   NextState = "game_end"
)

type NextResponse[T NextRoundImage | EndGameResult] struct {
	State NextState `json:"state"`
	Data  T         `json:"data"`
}

type NextRoundImage struct {
	ImageURI string `json:"image_uri"`
}

type EndGameResult struct {
	Result map[string]EndGame `json:"result"`
}

type EndGame struct {
	PerUser map[int]Game `json:"per_user"`
	Score   int          `json:"img_score"`
}

type Game struct {
	Username string `json:"username"`
	Img      string `json:"img"`
	Prompt   string `json:"answer"`
}

func findGame(data map[string][]model.InGamePrompt, conID string, index int) *model.InGamePrompt {
	for _, prompt := range data[conID] {
		if int(prompt.GameIndex) == index {
			return &prompt
		}
	}
	return nil
}

func (i *GameInteractor) NextRound(ctx context.Context, roomID string) error {
	DownAnsweredFlag(roomID)

	room, err := i.repo.GetRoom(ctx, roomID)
	if err != nil {
		return err
	}

	players, err := i.repo.GetConnectedPlayers(ctx, roomID)
	if err != nil {
		return err
	}

	if room.CurrentGame >= room.GameSize {
		prompts, err := i.repo.GetIngamePrompts(ctx, roomID)
		if err != nil {
			return err
		}
		// conID : name
		userMap := make(map[string]string, len(players))
		useriMap := make(map[int]string, len(players))
		for _, player := range players {
			userMap[player.ConnectionID] = player.Username
			useriMap[int(player.Index)] = player.ConnectionID
		}

		perGame := make(map[string][]model.InGamePrompt, room.GameSize)
		for _, prompt := range prompts {
			perGame[prompt.ConnectionID] = append(perGame[prompt.ConnectionID], prompt)
		}

		result := make(map[string]map[int]Game, room.GameSize)

		for i, conID := range useriMap {
			result[conID] = make(map[int]Game, room.GameSize)
			for j := 1; j <= int(room.GameSize); j++ {
				useri := j + i - 1
				if useri > len(players) {
					useri = useri - len(players)
				}
				inGame := findGame(perGame, useriMap[useri], j)
				if inGame != nil {
					result[conID][j] = Game{
						Username: userMap[inGame.ConnectionID],
						Img:      fmt.Sprintf("%s/%s/%s/%d.jpg", config.Config.Aws.CloudFrontURI, roomID, inGame.ConnectionID, j),
						Prompt:   inGame.Prompt,
					}
				}
			}
		}

		resultMap := EndGameResult{
			Result: make(map[string]EndGame, room.GameSize),
		}

		for conID, r := range result {
			result := EndGame{
				PerUser: r,
				Score:   0,
			}

			score, err := i.bedrock.ComparePrompt(ctx, r[0].Prompt, r[int(room.GameSize)-1].Prompt)
			if err != nil {
				return err
			}
			result.Score = score

			resultMap.Result[userMap[conID]] = result
		}

		data, err := json.Marshal(&NextResponse[EndGameResult]{
			State: StateGameEnd,
			Data:  resultMap,
		})
		if err != nil {
			return err
		}

		if err := dunnyBloadCast(roomID, StateGameEnd); err != nil {
			return err
		}

		if err := BroadcastInRoom(roomID, data); err != nil {
			return err
		}
		// remove room
		DeleteRoomSession(roomID)

		return nil
	}

	room.CurrentGame++
	err = i.repo.UpdateRoom(ctx, room)
	if err != nil {
		return err
	}

	if err := dunnyBloadCast(roomID, StateNextRound); err != nil {
		return err
	}

	playersMap := make(map[int]model.ConnectedPlayer, len(players))
	for _, player := range players {
		playersMap[int(player.Index)] = player
	}

	for _, player := range players {
		pindex := int(player.Index+room.CurrentGame) - 1
		if pindex > len(players) {
			pindex = pindex - len(players)
		}

		if err := sendImage(roomID, player.ConnectionID, fmt.Sprintf("%s/%s/%s/%d.jpg", config.Config.Aws.CloudFrontURI, roomID, playersMap[pindex].ConnectionID, room.CurrentGame-1)); err != nil {
			return err
		}
	}

	return nil

}

func sendImage(roomID, connectionID string, imageURI string) error {
	data, err := json.Marshal(&NextResponse[NextRoundImage]{
		State: StateNextRound,
		Data: NextRoundImage{
			ImageURI: imageURI,
		},
	})
	if err != nil {
		return err
	}

	return SendMessageByID(roomID, connectionID, data)
}

func dunnyBloadCast(roomID string, state NextState) error {
	data, err := json.Marshal(&Dummy{
		State: state,
	})
	if err != nil {
		return err
	}

	return BroadcastInRoom(roomID, data)
}
