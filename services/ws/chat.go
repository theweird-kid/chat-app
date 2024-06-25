package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/theweird-kid/chat-app/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ChatService struct {
	hub *Hub
}

func NewChatService(h *Hub) *ChatService {
	return &ChatService{
		hub: h,
	}
}

type CreateRoom struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *ChatService) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoom
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadGateway, fmt.Sprintln("Error Creating Room"))
		return
	}

	//Store Room to the Hub
	c.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	utils.RespondWithJSON(w, http.StatusOK, fmt.Sprintln("Room created Successfully"))
}

func (c *ChatService) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error Upgrading Connection", err)
		return
	}

	roomID := chi.URLParam(r, "room_id")
	clientID := r.URL.Query().Get("user_id")
	userName := r.URL.Query().Get("user")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: userName,
	}

	m := &Message{
		Content:  "User " + cl.Username + " Joined The Room",
		RoomID:   cl.RoomID,
		Username: cl.Username,
	}

	log.Println(m)

	//Register the Client through Register Channel
	c.hub.Register <- cl

	//Broadcast the message
	c.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(c.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *ChatService) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]RoomRes, 0)

	for _, r := range c.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	utils.RespondWithJSON(w, http.StatusOK, rooms)
}

type ClientRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *ChatService) GetClients(w http.ResponseWriter, r *http.Request) {
	var clients []ClientRes
	roomID := chi.URLParam(r, "room_id")

	if _, ok := c.hub.Rooms[roomID]; !ok {
		clients = make([]ClientRes, 0)
		utils.RespondWithJSON(w, http.StatusOK, clients)
	}

	for _, client := range c.hub.Rooms[roomID].Clients {
		clients = append(clients, ClientRes{
			client.ID,
			client.Username,
		})
	}

	utils.RespondWithJSON(w, http.StatusOK, clients)
}
