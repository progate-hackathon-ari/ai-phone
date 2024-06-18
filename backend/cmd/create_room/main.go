package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/progate-hackathon-ari/backend/internal/handler"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), "failed to run", "error", err)
	}
}

func run() error {
	lambda.Start(handler.CreateRoom)
	return nil
}
