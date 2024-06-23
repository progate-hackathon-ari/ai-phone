package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/repository"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type RoomInteractor struct {
	repository repository.DataAccess
}

func NewRoomInteractor(repository repository.DataAccess) *RoomInteractor {
	return &RoomInteractor{
		repository: repository,
	}
}

func (i *RoomInteractor) CreateRoom(ctx context.Context) (*model.Room, error) {

	roomID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create room id"))
	}

	room := &model.Room{
		RoomID:      roomID.String(),
		IsStarted:   false,
		GameSize:    1,
		CurrentGame: 0,
	}

	if err := i.repository.CreateRoom(ctx, room); err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create room"))
	}

	return room, nil
}

func (i *RoomInteractor) UpdateRoom(ctx context.Context, roomID, extraPrompt string) (*model.Room, error) {
	if extraPrompt == "" {
		extraPrompt = "\n"
	}
	log.Println(extraPrompt)
	room := &model.Room{
		RoomID:      roomID,
		ExtraPrompt: extraPrompt,
		GameSize:    -1,
		CurrentGame: -1,
	}

	if err := i.repository.UpdateRoom(ctx, room); err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to update room"))
	}

	return room, nil
}
