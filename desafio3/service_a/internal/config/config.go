package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	path_file_env := os.Getenv("PATH_ROOT")

	err := godotenv.Load(filepath.Join(path_file_env, ".env"))
	if err != nil {
		log.Fatal("Não foi possível carregar o arquivo .env de RabbitMQ: %v ", err)
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
