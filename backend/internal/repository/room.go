package repository

import (
	"context"

	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type Room interface {
	CreateRoom(ctx context.Context, room *model.Room) error
}
