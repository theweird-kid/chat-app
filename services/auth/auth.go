package auth

import (
	"context"
	"github.com/theweird-kid/chat-app/config"

	"github.com/theweird-kid/chat-app/internal/database"

	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/theweird-kid/chat-app/utils"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const UserKey contextKey = "userID"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hash), err
}

func ComparePassword(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}

func MiddlewareAuth(handler http.HandlerFunc, db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetTokenFromRequest(r)
		token, err := ValidateToken(tokenString)
		if err != nil {
			log.Printf("Failed to validate token: %v", err)
			utils.RespondWithError(w, 300, fmt.Sprintln("Permission denied"))
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			utils.RespondWithError(w, 300, fmt.Sprintln("Permission denied"))
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)
		id, err := uuid.Parse(str)
		if err != nil {
			log.Printf("Failed to convert string to uuid: %v", err)
			utils.RespondWithError(w, 300, fmt.Sprintln("Server Problem"))
			return
		}

		u, err := db.GetUserByID(r.Context(), id)
		if err != nil {
			log.Println("failed to get user by ID:", err)
			utils.RespondWithError(w, 300, fmt.Sprintln("Permission denied"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handler(w, r)
	}
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth, err := r.Cookie("Authorization")
	if err != nil {
		return ""
	}
	if tokenAuth.Value != "" {
		return tokenAuth.Value
	}
	return ""
}

func ValidateToken(t string) (*jwt.Token, error) {
	log.Println(t) //debug
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.Secret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) uuid.UUID {
	userID, ok := ctx.Value(UserKey).(uuid.UUID)
	if !ok {
		return uuid.Nil
	}

	return userID
}
