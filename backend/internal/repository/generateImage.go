package repository

import (
	"context"
)

type GenerateImage interface {
	GenerateImage(ctx context.Context, roomId string, connectionId string, prompt string) error
}
