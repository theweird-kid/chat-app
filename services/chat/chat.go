package chat

import (
	"context"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type ChatService struct {
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (chat *ChatService) Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket endpoint hit")
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading connection to websocket", err)
	}
	defer conn.CloseNow()

	var v interface{}
	for {
		// Create a new context with timeout for each message read to reset the timeout
		ctx, cancel := context.WithTimeout(r.Context(), time.Minute*10)

		err = wsjson.Read(ctx, conn, &v)
		cancel() // Cancel the context as soon as the read is done or has failed

		if err != nil {
			// Log the error and break out of the loop if the error is critical
			log.Println("error reading json:", err)
			// Determine if the error is critical (e.g., connection closed) and break if so
			if websocket.CloseStatus(err) != websocket.StatusNormalClosure {
				break
			}
		} else {
			log.Printf("received: %v", v)
		}
	}
}
