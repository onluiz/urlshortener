package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/teris-io/shortid"
)

var db *sqlx.DB

type URLService struct {
	ShortenURL   func(longURL string) (string, error)
	GetURLByCode func(code string) (string, error)
}

type URL struct {
	Id        string  `db:"id"`
	Original  string  `db:"original"`
	Short     string  `db:"short"`
	Code      string  `db:"code"`
	CreatedAt []uint8 `db:"created_at"`
}

func NewURLService(loggerInstance *log.Logger, dbInstance *sqlx.DB) *URLService {
	urlService := new(URLService)
	urlService.ShortenURL = shortenURL
	urlService.GetURLByCode = getURLByCode
	db = dbInstance
	return urlService
}

func shortenURL(longURL string) (string, error) {
	u, err := url.ParseRequestURI(longURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.String())

	code, err := shortid.Generate()
	if err != nil {
		panic(err)
	}
	db.MustExec(
		"INSERT INTO url (original, short, code, created_at) VALUES (?, ?, ?, ?)",
		u.String(),
		u.String(),
		code,
		time.Now(),
	)
	return code, nil
}

func getURLByCode(code string) (string, error) {
	url := URL{}
	err := db.Get(&url, "SELECT * FROM url url WHERE code=?", code)
	return url.Original, err
}
