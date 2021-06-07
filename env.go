package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Globals struct {
	MigrationsPath string

	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBSchema   string
	DBUri      string

	ServerHost    string
	ServerPort    string
	ServerBaseURL string
}

func NewEnv(logger *log.Logger) *Globals {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
		os.Exit(3)
	}

	globals := new(Globals)
	globals.MigrationsPath = os.Getenv("MIGRATIONS_PATH")

	globals.DBUser = os.Getenv("DB_USER")
	globals.DBPassword = os.Getenv("DB_PASSWORD")
	globals.DBHost = os.Getenv("DB_HOST")
	globals.DBPort = os.Getenv("DB_PORT")
	globals.DBSchema = os.Getenv("DB_SCHEMA")
	globals.DBUri = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_SCHEMA")

	globals.ServerHost = os.Getenv("SERVER_HOST")
	globals.ServerPort = os.Getenv("SERVER_PORT")
	globals.ServerBaseURL = globals.ServerHost + ":" + globals.ServerPort

	return globals
}
