package main

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv(logger *log.Logger) {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
}
