package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, logger *log.Logger, globals *Globals) chi.Router {
	var r chi.Router = chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go http.ListenAndServe(":"+globals.ServerPort, r)
			return nil
		},
	})

	return r
}
