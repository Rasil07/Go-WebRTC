package routes

import (
	"go_webRTC/api/controllers"
	"go_webRTC/infrastructure"
)

type ChatRoutes struct{
	chatController *controllers.ChatController
	server *infrastructure.Server
}

func NewChatRoutes(ctC *controllers.ChatController, svr *infrastructure.Server) *ChatRoutes{
	return &ChatRoutes{
		chatController: ctC,
		server: svr,
	}
}

func (cR *ChatRoutes) Setup(){
	routes:= cR.server.Group("/api/chat")

	routes.GET("/ping",cR.chatController.Pong)
}