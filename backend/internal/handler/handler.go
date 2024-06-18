package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayProxyHandler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func Response(status int, response any) (events.APIGatewayProxyResponse, error) {
	data, err := json.Marshal(response)
	if err != nil {
		return ErrResponse(http.StatusInternalServerError, err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		Body:            base64.StdEncoding.EncodeToString(data),
		IsBase64Encoded: true,
	}, nil
}

func ErrResponse(status int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       message,
	}, nil
}

func Unmarshal[T any](body string) (T, error) {
	var data T
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
