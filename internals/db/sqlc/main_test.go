package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sulabhghimire/bank-api/internals/config"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig("../../../")
	if err != nil {
		log.Fatal("Can't load config", err)
	}
	testDB, err = sql.Open(cfg.DB_DRIVER, cfg.DB_SOURCE)
	if err != nil {
		log.Fatal("Failed to connection to DB")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
