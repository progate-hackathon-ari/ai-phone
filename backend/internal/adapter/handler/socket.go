package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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

			msg = bytes.ReplaceAll(msg, []byte("\""), []byte(""))
			// decode base64

			message, err := unmarshal[RequestMessage](string(msg))
			if err != nil {
				log.Error(ctx, "failed to unmarshal or decode message", "error", err)
			}

			switch message.Event {
			case EventJoin:
				join, err := unmarshal[Join](message.Data)
				if err != nil {
					log.Error(ctx, "failed to unmarshal or decode message", "error", err)
					return err
				}

				if err := game.JoinRoom(ctx, message.RoomID, join.Name); err != nil {
					log.Error(ctx, "failed to join room", "error", err)
					// ゲームに入れないなら切断
					if !errors.Is(err, aerror.ErrFaliedMarshal) {
						return err
					}
				}

			case EventAnswer:
				answer, err := unmarshal[Answer](message.Data)
				if err != nil {
					log.Error(ctx, "failed to unmarshal or decode message", "error", err)
					return err
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
				count, err := unmarshal[CountDown](message.Data)
				if err != nil {
					log.Error(ctx, "failed to unmarshal or decode message", "error", err)
					return err
				}

				if err := game.CountDown(ctx, count.Count); err != nil {
					log.Error(ctx, "failed to count down", "error", err)
				}
			}
		}
	}
}

func unmarshal[T any](s string) (T, error) {
	msg, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return *new(T), err
	}
	fmt.Println(string(msg))

	str, err := url.QueryUnescape(string(msg))
	if err != nil {
		return *new(T), err
	}
	var t T
	if err = json.Unmarshal([]byte(str), &t); err != nil {
		return *new(T), err
	}

	return t, nil
}
