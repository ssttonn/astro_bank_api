package main

import (
	"astrobank/api"
	"database/sql"
	"log"

	db "astrobank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://astrotify-admin:sample@localhost:5432/astrobank_db?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, error := sql.Open(dbDriver, dbSource)
	if error != nil {
		log.Fatal("Can't connect to db: ", error)
	}

	store := db.NewStore(conn)
	server := api.NewServer(
		store,
	)

	if error = server.Start(serverAddress); error != nil {
		log.Fatal("Can't start server: ", error)
	}

}
