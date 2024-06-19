package usecase

import (
	"github.com/gorilla/websocket"
	"github.com/progate-hackathon-ari/backend/internal/repository"
)

type GameInteractor struct {
	repo   repository.DataAccess
	client *Client
}

func NewGameInteractor(ws *websocket.Conn, repo repository.DataAccess) *GameInteractor {
	return &GameInteractor{
		repo: repo,
		client: &Client{
			ws: ws,
		},
	}
}
