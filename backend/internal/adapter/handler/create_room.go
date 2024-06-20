package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

type CreateRoomRequest struct {
	ExtraPrompt string `json:"extraPrompt"`
}

type CreateRoomResponse struct {
	RoomID      string `json:"roomId"`
	ExtraPrompt string `json:"extraPrompt"`
}

func CreateRoom(i *usecase.CreateRoomInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var reqBody CreateRoomRequest
		if err := c.Bind(&reqBody); err != nil {
			log.Error(ctx, "failed to bind request", err.Error())
			return echo.ErrBadRequest
		}

		room, err := i.CreateRoom(ctx, reqBody.ExtraPrompt)
		if err != nil {
			log.Error(ctx, "faled to create room", err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, CreateRoomResponse{
			RoomID:      room.RoomID,
			ExtraPrompt: room.ExtraPrompt,
		})
	}
}
