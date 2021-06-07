package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ShortenURLPayload struct {
	LongURL string `json:"longURL"`
}

func RegisterRoutes(r chi.Router) {
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
		shortenedURL, error := shortenURL(data.LongURL)
		if error != nil {
			panic(err)
		}

		rw.Write([]byte(shortenedURL))
	})

	r.Get("/{code}", func(rw http.ResponseWriter, r *http.Request) {
		if code := chi.URLParam(r, "code"); code != "" {
			rw.Write([]byte(getURLByCode(code)))
		} else {
			rw.Write([]byte("Your code is not valid"))
		}
	})
}
