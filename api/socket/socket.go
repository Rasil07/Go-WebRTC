package socket

import (
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
)




type Websocket struct{
	 *websocket.Upgrader
}

func NewSocket() *Websocket{
	return &Websocket{
		&websocket.Upgrader{},
	}
}

var Module = fx.Options(
	fx.Provide(NewSocket),
)