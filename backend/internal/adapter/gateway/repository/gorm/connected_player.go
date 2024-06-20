package gorm

import (
	"context"

	"github.com/progate-hackathon-ari/backend/internal/entities/model"
	"gorm.io/gorm"
)

type ConnectedPlayerRepository struct {
	db *gorm.DB
}

func NewConnectedPlayerRepository(db *gorm.DB) *ConnectedPlayerRepository {
	return &ConnectedPlayerRepository{db: db}
}

func (r *ConnectedPlayerRepository) CreateConnectedPlayer(ctx context.Context, param model.ConnectedPlayer) error {
	return r.db.Create(&param).Error
}
func (r *ConnectedPlayerRepository) GetConnectedPlayers(ctx context.Context, roomID string) ([]model.ConnectedPlayer, error) {
	var players []model.ConnectedPlayer
	err := r.db.Where("room_id = ?", roomID).Find(&players).Error
	return players, err
}
