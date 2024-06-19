package usecase

import (
	"github.com/gorilla/websocket"
	"github.com/progate-hackathon-ari/backend/internal/external/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/external/s3"
	"github.com/progate-hackathon-ari/backend/internal/repository"
)

type GameInteractor struct {
	repo    repository.DataAccess
	s3      s3.S3
	bedrock bedrock.Bedrock
	client  *Client
}

func NewGameInteractor(ws *websocket.Conn, repo repository.DataAccess, s3 s3.S3, bedrock bedrock.Bedrock) *GameInteractor {
	return &GameInteractor{
		repo:    repo,
		s3:      s3,
		bedrock: bedrock,
		client: &Client{
			ws: ws,
		},
	}
}
