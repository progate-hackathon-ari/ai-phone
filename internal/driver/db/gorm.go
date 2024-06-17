package db

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGORM(db *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return gormDB
}
