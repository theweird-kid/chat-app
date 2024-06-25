package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/theweird-kid/chat-app/internal/database"
	"github.com/theweird-kid/chat-app/services/user"
	"github.com/theweird-kid/chat-app/services/ws"
)

type Service interface {
	RegisterRoutes()
}
type APIServer struct {
	addr string
	db   *sql.DB
	q    *database.Queries
}

func NewAPIServer(addr string, db *sql.DB, q *database.Queries) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
		q:    q,
	}
}

func (app *APIServer) RunServer() error {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 300,
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

	//chat Service
	hub := ws.NewHub()
	chatService := ws.NewChatService(hub)
	chatService.RegisterRoutes(router)
	go hub.Run()

	log.Println("Starting server on port", app.addr)
	return http.ListenAndServe(app.addr, router)

}
