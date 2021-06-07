package main

import (
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewDB,
			NewURLService,
			NewServer,
		),
		fx.Invoke(InitEnv),
		fx.Invoke(RegisterRoutes),
	)
	app.Run()
}
