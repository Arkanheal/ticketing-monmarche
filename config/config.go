package config

import (
	"github.com/joho/godotenv"

	"os"
)

func Config(key string) string {

	err := godotenv.Load(".env") // .env.prod .env.dev

	if err != nil {
		panic("Erreur dans le chargement du .env")
	}

	return os.Getenv(key)
}
