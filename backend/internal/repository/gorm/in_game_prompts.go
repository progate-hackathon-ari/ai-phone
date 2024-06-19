package gorm

import (
	"context"

	"github.com/progate-hackathon-ari/backend/internal/entities/model"
	"gorm.io/gorm"
)

type IngamePromptRepository struct {
	db *gorm.DB
}

func NewIngamePromptRepository(db *gorm.DB) *IngamePromptRepository {
	return &IngamePromptRepository{db: db}
}

func (r *IngamePromptRepository) CreateIngamePrompt(ctx context.Context, room *model.InGamePrompt) error {
	return r.db.Create(room).Error
}

func (r *IngamePromptRepository) GetIngamePrompts(ctx context.Context, roomID string) ([]model.InGamePrompt, error) {
	var prompts []model.InGamePrompt
	err := r.db.Find(&prompts).Error
	return prompts, err
}
