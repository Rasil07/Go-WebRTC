package controllers

import (
	"go.uber.org/fx"
)


var Moudle = fx.Options(
	fx.Provide(NewChatController),
)


