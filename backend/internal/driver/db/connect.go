package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/progate-hackathon-ari/backend/cmd/config"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

func Connect() *sql.DB {
	ctx := context.Background()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name,
	)

	var db *sql.DB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(ctx, "Error parsing config", "error", err)
	}

	const maxRetries = 5
	const retryDelay = 2

	for i := 1; i <= maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		log.Warn(ctx, fmt.Sprintf("Error pinging DB (Attempt %d/%d) dsn:%s err:%s\n", i, maxRetries, dsn, err))

		if i < maxRetries {
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		log.Fatal(ctx, "Exceeded maximum retries: Error pinging DB", "error", err)
	}

	log.Info(ctx, "Connected to database")
	return db
}
