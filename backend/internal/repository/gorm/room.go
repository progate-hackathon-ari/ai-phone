package gorm

import (
	"context"
	"fmt"

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

func (r *RoomRepository) UpdateRoom(ctx context.Context, room *model.Room) error {
	if room.RoomID == "" {
		return fmt.Errorf("room id is empty")
	}

	updateRoom := make(map[string]any)
	updateRoom["room_id"] = room.RoomID

	if room.ExtraPrompt != "" {
		updateRoom["extra_prompt"] = room.ExtraPrompt
	}

	if room.GameSize > -1 {
		updateRoom["game_size"] = room.GameSize
	}

	if room.CurrentGame > -1 {
		updateRoom["current_game"] = room.CurrentGame
	}

	return r.db.Model(room).Updates(updateRoom).Error
}

func (r *RoomRepository) GetRoom(ctx context.Context, roomID string) (*model.Room, error) {
	var room model.Room
	if err := r.db.Where("room_id = ?", roomID).First(&room).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *RoomRepository) StartGame(ctx context.Context, roomID string) error {
	updatedRoom := map[string]any{
		"room_id":    roomID,
		"is_started": true,
	}

	return r.db.Model(&model.Room{}).Where("room_id = ?", roomID).Updates(updatedRoom).Error
}

var _ repository.Room = (*RoomRepository)(nil)
