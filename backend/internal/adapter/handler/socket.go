package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/repository"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/s3"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

var (
	upgrader = websocket.Upgrader{}
)

type Event string

const (
	EventJoin   Event = "join"
	EventAnswer Event = "answer"
	EventReady  Event = "ready"
	EventNext   Event = "next"
)

type RequestMessage struct {
	Event  Event  `json:"event"`
	RoomID string `json:"room_id"`
	Data   string `json:"data"`
}

type Join struct {
	Name string `json:"name"`
}

type Answer struct {
	Answer string `json:"answer"`
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

			// Write
			err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			_, msg, err := ws.ReadMessage()
			if err != nil {
				c.Logger().Error(err)
			}

			var message RequestMessage
			if err = json.Unmarshal(msg, &message); err != nil {
				c.Logger().Error(err)
			}
			log.Info(ctx, "message", message)

			switch message.Event {
			case EventJoin:
				var join Join
				if err = json.Unmarshal([]byte(message.Data), &join); err != nil {
					c.Logger().Error(err)
				}

				log.Info(ctx, "join", join)

				result, err := game.JoinRoom(ctx, message.RoomID, join.Name)
				if err != nil {
					c.Logger().Error(err)
				}

				data, err := json.Marshal(result)
				if err != nil {
					c.Logger().Error(err)
				}

				if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
					c.Logger().Error(err)
				}

			case EventAnswer:
				var answer Answer
				if err = json.Unmarshal([]byte(message.Data), &answer); err != nil {
					c.Logger().Error(err)
				}

				if err := game.ImageGenerate(ctx, message.RoomID, answer.Answer); err != nil {
					c.Logger().Error(err)
				}

			case EventReady:
				if err := game.ReadyGame(ctx, message.RoomID); err != nil {
					c.Logger().Error(err)
				}
			case EventNext:
				if err := game.NextRound(ctx, message.RoomID); err != nil {
					c.Logger().Error(err)
				}
			}
		}
	}
}
