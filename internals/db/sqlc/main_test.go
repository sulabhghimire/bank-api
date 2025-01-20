package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dBSource = "postgresql://admin:admin@localhost:5432/simple_bank"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dBSource)
	if err != nil {
		log.Fatal("Failed to connection to DB")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
