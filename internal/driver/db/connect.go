package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/XSAM/otelsql"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/progate-hackathon-ari/backend/cmd/config"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

func Connect() *sql.DB {
	ctx := context.Background()

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d database=%s sslmode=disable",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name,
	)

	var db *sql.DB
	pgxConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(ctx, "Error parsing config", "error", err)
	}
	conn := stdlib.GetConnector(*pgxConfig)
	if config.Config.Otel.IsUse {
		db = otelsql.OpenDB(conn,
			otelsql.WithSpanOptions(otelsql.SpanOptions{
				Ping:                 false,
				OmitConnResetSession: true,
				OmitConnPrepare:      true,
				OmitRows:             true,
				OmitConnectorConnect: true,
			}))
	} else {
		db = sql.OpenDB(conn)
	}

	const maxRetries = 5
	const retryDelay = 2

	for i := 1; i <= maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		log.Warn(ctx, fmt.Sprintf("Error pinging DB (Attempt %d/%d): %s\n", i, maxRetries, err))

		if i < maxRetries {
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		log.Fatal(ctx, "Exceeded maximum retries: Error pinging DB", "error", err)
	}

	return db
}
