package auth

import (
	"time"

	"github.com/theweird-kid/chat-app/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateJWT(secret []byte, userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"expireAt": time.Now().Add(time.Minute * time.Duration(config.Envs.TokenExpiration)).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
