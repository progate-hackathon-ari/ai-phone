package handler

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
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

func CreateRoom(i *usecase.CreateRoomInteractor) APIGatewayProxyHandler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Info(ctx, "CreateRoom", request.HTTPMethod, request.Path, request.Body)
		if request.HTTPMethod != "POST" {
			return ErrResponse(http.StatusBadRequest, "invalid method")
		}

		reqBody, err := Unmarshal[CreateRoomRequest](request.Body)
		if err != nil {
			return ErrResponse(http.StatusBadRequest, err.Error())
		}

		room, err := i.CreateRoom(ctx, reqBody.ExtraPrompt)
		if err != nil {
			return ErrResponse(http.StatusInternalServerError, err.Error())
		}

		return Response(http.StatusOK, CreateRoomResponse{
			RoomID:      room.RoomID,
			ExtraPrompt: room.ExtraPrompt,
		})
	}
}
