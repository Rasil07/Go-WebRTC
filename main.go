package main

import (
	"context"
	"go_webRTC/api/routes"
	"go_webRTC/appRegister"
	"go_webRTC/infrastructure"

	"go.uber.org/fx"
)



func main(){
		app:= fx.New(
			appRegister.Module,
			fx.Invoke(registerHooks),
		)
		app.Run()
}


func registerHooks(
	lifecycle fx.Lifecycle,
	ser *infrastructure.Server,
	rts routes.Routes,
){
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				rts.Setup()
				go ser.Run()
				return nil
			},
		},
	)
}