package db

import (
	"database/sql"
	"embed"
	"log"

	sqlc "github.com/cvzamannow/service-auth/internal/db/sqlc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migration/*.sql
var fs embed.FS

func Migrate(source string) {
	d, err := iofs.New(fs, "migration")

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, source)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()

}

func NewDBConnection(driver string, source string) *sqlc.Queries {
	db, err := sql.Open(driver, source)

	if err != nil {
		panic(err)
	}

	defer Migrate(source)

	q := sqlc.New(db)
	return q
}
