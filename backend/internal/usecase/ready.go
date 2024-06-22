package usecase

import (
	"context"
	"fmt"
)

func (i *GameInteractor) ReadyGame(ctx context.Context, roomID string) error {
	room, err := i.repo.GetRoom(ctx, roomID)
	if err != nil {
		return err
	}

	if room.IsStarted {
		return fmt.Errorf("room %s is already started", roomID)
	}

	if !IsMaster(roomID, i.client.info.ConnectionID) {
		return fmt.Errorf("you are not master")
	}

	if err := i.repo.StartGame(ctx, roomID); err != nil {
		return err
	}

	if err := dunnyBloadCast(roomID, StateStartGame); err != nil {
		return err
	}

	Counter(roomID, 3)

	return nil
}
