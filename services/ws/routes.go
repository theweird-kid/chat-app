package ws

import "github.com/go-chi/chi/v5"

func (c *ChatService) RegisterRoutes(r *chi.Mux) {
	r.Post("/ws/create_room", c.CreateRoom)
	r.Get("/ws/join_room/{room_id}", c.JoinRoom)
	r.Get("/ws/get_rooms", c.GetRooms)
	r.Get("/ws/get_clients/{room_id}", c.GetClients)
}
