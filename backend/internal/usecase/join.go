package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type JoinRoomResult struct {
	ConnectionID string                  `json:"connection_id"`
	Players      []model.ConnectedPlayer `json:"players"`
}

func (i *GameInteractor) JoinRoom(ctx context.Context, roomID, name string) (*JoinRoomResult, error) {
	players, err := i.repo.GetConnectedPlayers(ctx, roomID)
	if err != nil {
		return nil, err
	}

	room, err := i.repo.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	if room.IsStarted {
		return nil, fmt.Errorf("game started")
	}

	user := model.ConnectedPlayer{
		RoomID:       roomID,
		ConnectionID: uuid.NewString(),
		Index:        int32(len(players) + 1),
		Username:     name,
	}

	if err := i.repo.CreateConnectedPlayer(ctx, user); err != nil {
		return nil, err
	}

	i.client.info = &user

	players = append(players, user)

	var gamesize int
	if len(players) > 5 {
		gamesize = 5
	} else {
		gamesize = len(players)
	}

	if err := i.repo.UpdateRoom(ctx, &model.Room{
		RoomID:   roomID,
		GameSize: int32(gamesize),
	}); err != nil {
		return nil, err
	}

	AddClient(i.client.ws, &user, roomID)

	if err := BroadcastInRoom(roomID, []byte(fmt.Sprintf("player %d joined", user.Index))); err != nil {
		return nil, err
	}

	return &JoinRoomResult{
		ConnectionID: user.ConnectionID,
		Players:      players,
	}, nil
}
