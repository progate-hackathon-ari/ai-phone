package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/progate-hackathon-ari/backend/cmd/config"
	"github.com/progate-hackathon-ari/backend/internal/container"
	"github.com/progate-hackathon-ari/backend/internal/handler"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

func init() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(context.Background(), "failed to load env", "error", err)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), "failed to run", "error", err)
	}
}

func run() error {
	if err := container.NewContainer(); err != nil {
		return errors.Join(err, errors.New("failed to create container"))
	}

	sqlDB := container.Invoke[*sql.DB]()
	defer sqlDB.Close()
	roomInteractor := container.Invoke[*usecase.CreateRoomInteractor]()

	lambda.Start(handler.CreateRoom(roomInteractor))
	return nil
}
