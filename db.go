package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func NewDB(lc fx.Lifecycle, logger *log.Logger, globals *Globals) *sqlx.DB {
	db, err := sqlx.Connect("mysql", globals.DBUri)
	if err != nil {
		logger.Fatalln(err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	logger.Println("Database connected")

	return db
}
