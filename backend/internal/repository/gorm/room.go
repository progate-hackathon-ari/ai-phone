package gorm

import (
	"context"

	"github.com/progate-hackathon-ari/backend/internal/entities/model"
	"github.com/progate-hackathon-ari/backend/internal/repository"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) CreateRoom(ctx context.Context, room *model.Room) error {
	return r.db.Create(room).Error
}

var _ repository.Room = (*RoomRepository)(nil)
