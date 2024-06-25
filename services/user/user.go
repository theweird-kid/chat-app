package user

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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

}
