package config

import (
	"github.com/joho/godotenv"

	"os"
)

func Config(key string) string {

	err := godotenv.Load(".env") // .env.prod .env.dev

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
