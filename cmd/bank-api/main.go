package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sulabhghimire/bank-api/internals/api"
	"github.com/sulabhghimire/bank-api/internals/config"
	db "github.com/sulabhghimire/bank-api/internals/db/sqlc"
)

func main() {
	cfg, err := config.LoadConfig("./../../")
	if err != nil {
		log.Fatal("Can't load config variables")
	}

	conn, err := sql.Open(cfg.DB_DRIVER, cfg.DB_SOURCE)
	if err != nil {
		log.Fatal("Failed to connection to DB")
	}
	log.Println("Database connection successful")

	store := db.NewStore(conn)

	server := api.NewServer(store)
	err = server.Start(cfg.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}
