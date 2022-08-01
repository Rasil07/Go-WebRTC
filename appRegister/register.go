package appRegister

import (
	"go_webRTC/api/controllers"
	"go_webRTC/api/routes"
	"go_webRTC/infrastructure"

	"go.uber.org/fx"
)

var Module = fx.Options(
	routes.Module,
	controllers.Moudle,
	infrastructure.Module,
)

