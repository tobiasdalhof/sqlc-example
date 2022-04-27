package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var queries *Queries

func TestMain(m *testing.M) {
	setupDatabase()
	code := m.Run()
	os.Exit(code)
}

func setupDatabase() {
	db, err := sql.Open("postgres", "host=postgres-test user=default password=secret dbname=default sslmode=disable")
	if err != nil {
		log.Fatalf("could not open database connection: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("could not ping database: %s", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create postgres migration driver: %s", err)
	}

	path, _ := os.Getwd()
	if err != nil {
		log.Fatalf("could not get working directory: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s/../migrations", path), "postgres", driver)
	if err != nil {
		log.Fatalf("could not create migration db instance: %s", err)
	}

	err = m.Down()
	if err != nil {
		log.Fatalf("migrate down failed: %s", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("migrate up failed: %s", err)
	}

	queries = New(db)
}
