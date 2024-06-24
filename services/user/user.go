package user

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/theweird-kid/chat-app/services/chat"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) RegisterRoutes(router *chi.Mux) {
	router.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		log.Println("User route hit")
		w.Write([]byte("User Service"))
	})

	//chat Service
	chatService := chat.NewChatService()

	router.Get("/chat", chatService.Handler)
}
