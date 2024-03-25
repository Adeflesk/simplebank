package main

import (
	"database/sql"
	"log"

	db "github.com/Adeflesk/simplebank/db/sqlc/db/sqlc"
	"github.com/techschool/simplebank/api"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
