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
	dBSource = "postgresql://admin:admin@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dBSource)
	if err != nil {
		log.Fatal("Failed to connection to DB")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
