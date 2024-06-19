package handler

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

type GenerateImageRequest struct {
	RoomId       string `json:roomId`
	ConnectionId string `json:connectionId`
	Prompt       string `json:prompt`
}

type GenerateImageResponse struct {
	status string `json:status`
}

func GenerateImage(i *usecase.GenerateImageInteractor) APIGatewayProxyHandler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Info(ctx, "GenerateImage", request.HTTPMethod, request.Path, request.Body)

		if request.HTTPMethod != "POST" {
			return ErrResponse(http.StatusBadRequest, "invalid method")
		}

		reqBody, err := Unmarshal[GenerateImageRequest](request.Body)
		if err != nil {
			return ErrResponse(http.StatusBadRequest, err.Error())
		}

		result, err := i.GenerateImage(ctx, reqBody.RoomId, reqBody.ConnectionId, reqBody.Prompt)
		if err != nil {
			return ErrResponse(http.StatusInternalServerError, err.Error())
		}

		return Response(http.StatusOK, GenerateImageResponse{
			status: result,
		})
	}
}
