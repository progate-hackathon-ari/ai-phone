package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/progate-hackathon-ari/backend/cmd/config"
	"github.com/progate-hackathon-ari/backend/cmd/migrate/migration"
	"github.com/progate-hackathon-ari/backend/internal/driver/db"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

var migratefile string

type envFlag []string

func (e *envFlag) String() string {
	return strings.Join(*e, ",")
}

func (e *envFlag) Set(v string) error {
	*e = append(*e, v)
	return nil
}

func init() {
	var envFile envFlag
	flag.StringVar(&migratefile, "f", "", "migrate file path")
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
	db := db.Connect()
	defer db.Close()

	// migrate
	m, err := migration.Migrate(db, "file://"+migratefile, nil)
	if err != nil {
		log.Error(context.Background(), "migrate new error", "error", err)
		return fmt.Errorf("migrate new error: %w", err)
	}

	if len(os.Args) != 6 {
		log.Error(context.Background(), "invalid command")
		return fmt.Errorf("invalid command: %s", os.Args)
	}

	command := os.Args[5]

	switch command {
	case "up":
		// migrate up
		if err := m.Up(); err != nil {
			// 変更がない場合は無視
			if err != migrate.ErrNoChange {
				log.Error(context.Background(), "migrate up error", "error", err)
				return fmt.Errorf("'migrate up error: %w", err)
			}
		}
	case "down":
		// migrate down
		if err := m.Down(); err != nil {
			log.Error(context.Background(), "migrate down error", "error", err)
			return fmt.Errorf("'migrate down error: %w", err)
		}
	default:
		log.Error(context.Background(), "unknown command", "command", command)
		return fmt.Errorf("unknown command: %s", command)
	}
	log.Info(context.Background(), "migrate ' successful")
	return nil
}
