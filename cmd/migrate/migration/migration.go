package migration

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB, file string, arg any) (*migrate.Migrate, error) {
	if arg == nil {
		arg = &postgres.Config{}
	}
	return migratePostgres(db, file, arg.(*postgres.Config))
}
func migratePostgres(db *sql.DB, file string, arg *postgres.Config) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, arg)
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		file,
		"postgres",
		driver,
	)
}
