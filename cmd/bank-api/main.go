package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sulabhghimire/bank-api/internals/api"
	db "github.com/sulabhghimire/bank-api/internals/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dBSource      = "postgresql://admin:admin@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "localhost:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dBSource)
	if err != nil {
		log.Fatal("Failed to connection to DB")
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}
