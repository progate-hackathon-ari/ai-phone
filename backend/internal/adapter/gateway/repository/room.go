package repository

import (
	"context"

	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type Room interface {
	CreateRoom(ctx context.Context, room *model.Room) error
	GetRoom(ctx context.Context, roomID string) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
	StartGame(ctx context.Context, roomID string) error
}
