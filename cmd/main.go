package main

import (
	"github.com/theweird-kid/chat-app/cmd/api"
)

func main() {
	app := api.NewAPIServer(":8080")
	app.RunServer()
}
