package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func NewDB(lc fx.Lifecycle, logger *log.Logger) *sqlx.DB {
	dbUri := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_SCHEMA")
	db, err := sqlx.Connect("mysql", dbUri)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Database connected")

	return db
}
