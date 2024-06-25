package main

import (
	"log"

	"github.com/theweird-kid/chat-app/cmd/api"
	"github.com/theweird-kid/chat-app/config"
	"github.com/theweird-kid/chat-app/db"
)

func main() {
	db, q, err := db.NewPostgresStorage(config.Envs.DATABASE_URL)
	if err != nil {
		log.Fatal("Couldn't connect to the database")
	}
	app := api.NewAPIServer(config.Envs.PORT, db, q)
	app.RunServer()

}
