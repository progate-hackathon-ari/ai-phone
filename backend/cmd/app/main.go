package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"strings"

	"github.com/progate-hackathon-ari/backend/cmd/config"
	"github.com/progate-hackathon-ari/backend/internal/container"
	"github.com/progate-hackathon-ari/backend/internal/router"
	"github.com/progate-hackathon-ari/backend/internal/server"
)

type envFlag []string

func (e *envFlag) String() string {
	return strings.Join(*e, ",")
}

func (e *envFlag) Set(v string) error {
	*e = append(*e, v)
	return nil
}

func init() {
	// Usage: eg. go run main.go -e .env -e hoge.env -e fuga.env ...
	var envFile envFlag
	flag.Var(&envFile, "e", "path to .env file \n eg. -e .env -e another.env . ")
	flag.Parse()

	if err := config.LoadEnv(envFile...); err != nil {
		log.Fatal(context.Background(), "lod Env Error", "error", err)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), "failed to run", "error", err)
	}
}

func run() error {
	if err := container.NewContainer(); err != nil {
		return err
	}

	sqlDB := container.Invoke[*sql.DB]()
	defer sqlDB.Close()

	router := router.NewRouter()

	if err := server.New(":8080", router).RunWithGraceful(); err != nil {
		return err
	}

	return nil
}
