package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type ChatService struct {
}

func NewChatService() *ChatService {
	return &ChatService{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (chat *ChatService) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error connecting to the websocket", err)
	}
	defer conn.Close()

	go handleChat(conn)

}

func handleChat(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("error connecting to the websocket", err)
		}
		log.Println(p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
