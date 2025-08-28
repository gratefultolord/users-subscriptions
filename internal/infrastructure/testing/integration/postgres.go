package integration

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
)

type TestPostgres struct {
	DB       *sql.DB
	Resource *dockertest.Resource
	Pool     *dockertest.Pool
}

func NewPostgres(t *testing.T) *TestPostgres {
	t.Helper()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("dockertest.NewPool: %v", err)
	}

	resource, err := pool.Run("postgres", "15", []string{
		"POSTGRES_USER=test_user",
		"POSTGRES_PASSWORD=test_password",
		"POSTGRES_DB=test_db",
	})
	if err != nil {
		t.Fatalf("pool.Run: %v", err)
	}

	dsn := fmt.Sprintf("postgres://test_user:test_password@localhost:%s/test_db?sslmode=disable",
		resource.GetPort("5432/tcp"))

	var db *sql.DB

	if err := pool.Retry(func() error {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return fmt.Errorf("sql.Open: %w", err)
		}

		return db.Ping()
	}); err != nil {
		t.Fatalf("pool.Retry: %v", err)
	}

	return &TestPostgres{
		DB:       db,
		Resource: resource,
		Pool:     pool,
	}
}

func (pg *TestPostgres) CleanUp() {
	if err := pg.Pool.Purge(pg.Resource); err != nil {
		log.Printf("pg.Pool.Purge: %v", err)
	}
}

func (pg *TestPostgres) TruncateTables(t *testing.T, tables ...string) {
	t.Helper()

	for _, table := range tables {
		_, err := pg.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			t.Fatalf("pg.DB.Exec: %v", err)
		}
	}
}
