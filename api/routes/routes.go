package routes

import "go.uber.org/fx"



var Module = fx.Options(
	fx.Provide(NewChatRoutes),

	fx.Provide(NewRoutes),
)

type Routes []Route
type Route interface{
	Setup()
}

func NewRoutes(chatRoutes *ChatRoutes) Routes{
	return Routes{
		chatRoutes,
	}
}

func (r Routes) Setup(){
	for _, route:=range r{
		route.Setup()
	}
}

