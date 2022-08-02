package main

import (
	"gateway-service/client"
	"gateway-service/config"
	"gateway-service/controller"
	"gateway-service/lib"
	"gateway-service/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(start()).Run()
}

func start() fx.Option {
	return fx.Options(
		fx.Provide(config.NewConfig),
		lib.Module,
		client.Module,
		server.Module,
		controller.Module,
	)
}
