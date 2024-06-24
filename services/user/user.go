package user

import (
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
		w.Write([]byte("User Service"))
	})

	//chat Service
	chat := &chat.ChatService{}

	router.Get("/chat", chat.Handler)
}