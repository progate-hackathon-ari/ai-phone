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

type NextResponse[T NextRoundImage | EndGame] struct {
	State NextState `json:"state"`
	Data  T         `json:"data"`
}

type NextRoundImage struct {
	ImageURI string `json:"image_uri"`
}

type OneGame struct {
	Prompt   string `json:"prompt"`
	ImageURI string `json:"image_uri"`
}

type EndGame struct {
	Result map[string]map[int]OneGame `json:"result"`
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

		promptMap := make(map[string]map[int]string, room.GameSize)

		for _, prompt := range prompts {
			promptMap[prompt.ConnectionID] = make(map[int]string, room.GameSize)
			for i := 0; i < int(room.GameSize); i++ {
				promptMap[prompt.ConnectionID][i] = prompt.Prompt
			}
		}

		resultMap := make(map[string]map[int]OneGame, len(players))
		for _, player := range players {
			resultMap[player.ConnectionID] = make(map[int]OneGame, room.GameSize)
			for i := 0; i < int(room.GameSize); i++ {
				resultMap[player.ConnectionID][i] = OneGame{
					Prompt:   promptMap[player.ConnectionID][i],
					ImageURI: fmt.Sprintf("%s/%s/%s/%d.jpg", config.Config.Aws.CloudFrontURI, roomID, player.ConnectionID, i),
				}
			}
		}

		data, err := json.Marshal(&NextResponse[EndGame]{
			State: StateGameEnd,
			Data: EndGame{
				Result: resultMap,
			},
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
