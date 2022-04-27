package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
)

var queries *Queries

func TestMain(m *testing.M) {
	setupPostgres()
	code := m.Run()
	os.Exit(code)
}

func setupPostgres() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	_, err = pool.Run("postgres", "13", []string{
		"POSTGRES_DB=default",
		"POSTGRES_USER=default",
		"POSTGRES_PASSWORD=secret",
	})
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
	}

	pool.MaxWait = 10 * time.Second
	err = pool.Retry(func() error {
		db, err := sql.Open("postgres", "host=postgres user=default password=secret dbname=default sslmode=disable")
		if err != nil {
			return err
		}
		queries = New(db)
		return db.Ping()
	})

	if err != nil {
		log.Fatalf("could not connect to postgres container: %s", err)
	}
}
