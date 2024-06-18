package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
	"github.com/progate-hackathon-ari/backend/internal/repository"
)

type CreateRoomInteractor struct {
	repository repository.DataAccess
}

func NewCreateRoomInteractor(repository repository.DataAccess) *CreateRoomInteractor {
	return &CreateRoomInteractor{
		repository: repository,
	}
}

func (i *CreateRoomInteractor) CreateRoom(ctx context.Context, extraPrompts string) (*model.Room, error) {

	roomID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create room id"))
	}

	room := &model.Room{
		RoomID:      roomID.String(),
		ExtraPrompt: extraPrompts,
		IsStarted:   false,
		GameSize:    1,
		CurrentGame: 0,
	}

	if err := i.repository.CreateRoom(ctx, room); err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create room"))
	}

	return room, nil
}
