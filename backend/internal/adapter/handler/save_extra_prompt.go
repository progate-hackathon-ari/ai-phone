package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

type UpdateRoomRequest struct {
	RoomID      string `param:"room_id"`
	ExtraPrompt string `json:"extraPrompt"`
}

type UpdateRoomResponse struct {
	RoomID      string `json:"roomId"`
	ExtraPrompt string `json:"extraPrompt"`
}

func UpdateRoom(i *usecase.RoomInteractor) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var req UpdateRoomRequest
		if err := c.Bind(&req); err != nil {
			log.Error(ctx, "failed to bind request", err)
			return echo.ErrBadRequest
		}

		room, err := i.UpdateRoom(ctx, req.RoomID, req.ExtraPrompt)
		if err != nil {
			log.Error(ctx, "faled to create room", err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, UpdateRoomResponse{
			RoomID:      room.RoomID,
			ExtraPrompt: room.ExtraPrompt,
		})
	}
}
