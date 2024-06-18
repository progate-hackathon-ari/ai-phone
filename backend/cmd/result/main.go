package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), "failed to run", "error", err)
	}
}

func run() error {
	lambda.Start(requestHandler)
	return nil
}

func requestHandler(ctx context.Context) (string, error) {
	return "hello world", nil
}
