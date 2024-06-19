package usecase

import (
	"github.com/gorilla/websocket"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

type Client struct {
	ws   *websocket.Conn
	info *model.ConnectedPlayer
}
