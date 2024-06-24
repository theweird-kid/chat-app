package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/theweird-kid/chat-app/services/user"
)

type Service interface {
	RegisterRoutes()
}
type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (app *APIServer) RunServer() error {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	//subrouter
	subrouter := chi.NewRouter()
	router.Mount("/api/v1", subrouter)

	subrouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	//Services

	//user service
	user := user.NewUserService()
	user.RegisterRoutes(subrouter)

	log.Println("Starting server on port", app.addr)
	return http.ListenAndServe(app.addr, router)

}
