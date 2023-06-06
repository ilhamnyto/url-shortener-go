package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	POSTGRES_HOST = "POSTGRES_HOST"
	POSTGRES_PORT = "POSTGRES_PORT"
	POSTGRES_USER = "POSTGRES_USER"
	POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	POSTGRES_DB = "POSTGRES_DB"
	REDIS_HOST = "REDIS HOST"
	REDIS_PORT = "REDIS_PORT"
	REDIS_DB = "REDIS_DB"
	REDIS_PASSWORD = "REDIS_PASSWORD"
)

func LoadConfig(filename string) {
	godotenv.Load(filename)
}

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	str := os.Getenv(key)

	num, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return num
}