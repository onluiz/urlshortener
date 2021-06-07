package main

import (
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewEnv,
			NewDB,
			NewURLService,
			NewServer,
		),
		fx.Invoke(RegisterRoutes),
	)
	app.Run()
}
