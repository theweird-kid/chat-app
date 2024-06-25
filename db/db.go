package db

import (
	"database/sql"

	"github.com/theweird-kid/chat-app/internal/database"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	DB *database.Queries
}

func NewPostgresStorage(dbUrl string) (*sql.DB, *database.Queries, error) {
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, nil, err
	}
	q := database.New(conn)
	return conn, q, nil
}
