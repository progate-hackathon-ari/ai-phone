package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/repository"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/s3"
	"github.com/progate-hackathon-ari/backend/internal/aerror"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

var (
	upgrader = websocket.Upgrader{}
)

type Event string

const (
	EventJoin      Event = "join"
	EventAnswer    Event = "answer"
	EventReady     Event = "ready"
	EventNext      Event = "next"
	EventCountDown Event = "countdown"
)

type RequestMessage struct {
	Event  Event  `json:"event"`
	RoomID string `json:"roomId"`
	Data   string `json:"data"`
}

type Join struct {
	Name string `json:"name"`
}

type Answer struct {
	Answer string `json:"answer"`
}

type CountDown struct {
	Count int `json:"count"`
}

func SocketGameRoom(repo repository.DataAccess, s3 s3.S3, bedrock bedrock.Bedrock) echo.HandlerFunc {
	return func(c echo.Context) error {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		game := usecase.NewGameInteractor(ws, repo, s3, bedrock)

		ctx := context.Background()
		for {
			// Read
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Error(ctx, "falied to read message", "error", err)
				return err
			}

			var message RequestMessage
			if err = json.Unmarshal(msg, &message); err != nil {
				log.Error(ctx, "json unmarshal error", "error", err)
			}

			log.Debug(ctx, "message", message)

			switch message.Event {
			case EventJoin:
				var join Join
				if err = json.Unmarshal([]byte(message.Data), &join); err != nil {
					log.Error(ctx, "json unmarshal (2) error", "error", err)
				}

				if err := game.JoinRoom(ctx, message.RoomID, join.Name); err != nil {
					log.Error(ctx, "failed to join room", "error", err)
					// ゲームに入れないなら切断
					if !errors.Is(err, aerror.ErrFaliedMarshal) {
						return err
					}
				}

			case EventAnswer:
				var answer Answer
				if err = json.Unmarshal([]byte(message.Data), &answer); err != nil {
					log.Error(ctx, "json unmarshal (2) error", "error", err)
				}

				if err := game.ImageGenerate(ctx, message.RoomID, answer.Answer); err != nil {
					log.Error(ctx, "failed to image generate", "error", err)
				}

			case EventReady:
				if err := game.ReadyGame(ctx, message.RoomID); err != nil {
					log.Error(ctx, "failed to game start", "error", err)
				}

			case EventNext:
				if err := game.NextRound(ctx, message.RoomID); err != nil {
					log.Error(ctx, "failed to next round", "error", err)
				}
			case EventCountDown:
				var count CountDown
				if err = json.Unmarshal([]byte(message.Data), &count); err != nil {
					log.Error(ctx, "json unmarshal (2) error", "error", err)
				}

				if err := game.CountDown(ctx, count.Count); err != nil {
					log.Error(ctx, "failed to count down", "error", err)
				}
			}
		}
	}
}
