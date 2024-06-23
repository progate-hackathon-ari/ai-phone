package handler

import (
	"encoding/json"
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

		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}

		log.Info(ctx, "request", json_map["request"])

		c.Param("room_id")

		room, err := i.UpdateRoom(ctx, c.Param("room_id"), json_map["request"].(map[string]interface{})["extraPrompt"].(string))
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
