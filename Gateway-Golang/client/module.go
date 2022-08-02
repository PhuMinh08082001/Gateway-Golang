package client

import (
	"gateway-service/client/connect"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	connect.NewAccountCC,
	NewAccountClient,
)
