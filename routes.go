package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

type ShortenURLPayload struct {
	LongURL string `json:"longURL"`
}

func RegisterRoutes(r chi.Router, urlService *URLService, db *sqlx.DB, logger *log.Logger) {
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("You are at home"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/shorten", func(rw http.ResponseWriter, r *http.Request) {
		var data ShortenURLPayload
		err := render.DecodeJSON(r.Body, &data)
		if err != nil {
			panic(err)
		}
		shortenedURL, error := urlService.ShortenURL(data.LongURL)
		if error != nil {
			panic(err)
		}

		rw.Write([]byte(shortenedURL))
	})

	r.Get("/{code}", func(rw http.ResponseWriter, r *http.Request) {
		if code := chi.URLParam(r, "code"); code != "" {
			url, err := urlService.GetURLByCode(code)
			if err != nil {
				rw.Write([]byte("url not found"))
			} else {
				http.Redirect(rw, r, url, http.StatusMovedPermanently)
			}
		} else {
			logger.Panic("Code " + code + "is not valid")
			rw.Write([]byte("Your code is not valid"))
		}
	})
}
