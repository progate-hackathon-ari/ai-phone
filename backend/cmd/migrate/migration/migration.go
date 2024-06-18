package migration

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB, file string, arg any) (*migrate.Migrate, error) {
	if arg == nil {
		arg = &mysql.Config{}
	}
	return migrateMysql(db, file, arg.(*mysql.Config))
}
func migrateMysql(db *sql.DB, file string, arg *mysql.Config) (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(db, arg)
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		file,
		"mysql",
		driver,
	)
}
