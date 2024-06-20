package repository

import (
	"context"

	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type InGamePrompt interface {
	CreateIngamePrompt(ctx context.Context, room *model.InGamePrompt) error
	GetIngamePrompts(ctx context.Context, roomID string) ([]model.InGamePrompt, error)
}
