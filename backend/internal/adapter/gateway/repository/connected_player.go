package repository

import (
	"context"

	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type ConnectedPlayer interface {
	CreateConnectedPlayer(ctx context.Context, param model.ConnectedPlayer) error
	GetConnectedPlayers(ctx context.Context, roomID string) ([]model.ConnectedPlayer, error)
}
