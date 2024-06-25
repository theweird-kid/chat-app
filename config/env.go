package config

import (
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PORT            string
	DATABASE_URL    string
	TokenExpiration int64
	Secret          []byte
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load(".env")

	//get environment variables
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	expiration := os.Getenv("JWT_EXPIRATION")
	if expiration == "" {
		log.Fatal("JWT Expiration not found")
	}
	exp, err := strconv.ParseInt(expiration, 10, 64)
	if err != nil {
		log.Fatal("Problem with parsing expiry")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT secret not found")
	}

	return Config{
		PORT:            ":" + portString,
		DATABASE_URL:    dbURL,
		TokenExpiration: exp,
		Secret:          []byte(secret),
	}
}
