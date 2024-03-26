package main

import (
	"database/sql"
	"log"

	"github.com/Adeflesk/simplebank/api"
	db "github.com/Adeflesk/simplebank/db/sqlc"
	"github.com/Adeflesk/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// create a connection pool for our database
	conn, err := sql.Open(config.DBDRIVER, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
