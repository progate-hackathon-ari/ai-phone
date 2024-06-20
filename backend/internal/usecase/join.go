package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/progate-hackathon-ari/backend/internal/aerror"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type JoinRoomResult struct {
	ConnectionID string                  `json:"connection_id"`
	Players      []model.ConnectedPlayer `json:"players"`
}

func (i *GameInteractor) JoinRoom(ctx context.Context, roomID, name string) error {
	players, err := i.repo.GetConnectedPlayers(ctx, roomID)
	if err != nil {
		return err
	}

	room, err := i.repo.GetRoom(ctx, roomID)
	if err != nil {
		return err
	}

	if room.IsStarted {
		return fmt.Errorf("game started")
	}

	user := model.ConnectedPlayer{
		RoomID:       roomID,
		ConnectionID: uuid.NewString(),
		Index:        int32(len(players) + 1),
		Username:     name,
	}

	if err := i.repo.CreateConnectedPlayer(ctx, user); err != nil {
		return err
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
		return err
	}

	AddClient(i.client.ws, &user, roomID)

	data, err := json.Marshal(&JoinRoomResult{
		ConnectionID: user.ConnectionID,
		Players:      players,
	})
	if err != nil {
		return aerror.ErrFaliedMarshal
	}

	return BroadcastInRoom(roomID, data)
}
