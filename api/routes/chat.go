package routes

import (
	"fmt"
	"go_webRTC/api/controllers"
	"go_webRTC/api/socket"

	"go_webRTC/infrastructure"

	"github.com/gin-gonic/gin"
)

type ChatRoutes struct{
	chatController *controllers.ChatController
	server *infrastructure.Server
	socket *socket.Websocket

}

func NewChatRoutes(ctC *controllers.ChatController,skt *socket.Websocket, svr *infrastructure.Server) *ChatRoutes{
	return &ChatRoutes{
		chatController: ctC,
		server: svr,
		socket: skt,
	}
}

func (cR *ChatRoutes) Setup(){
	routes:= cR.server.Group("/api/chat")

	routes.GET("/ping",cR.chatController.Pong)
	routes.GET("/message",func(c *gin.Context){
		ws,err:= cR.socket.Upgrade(c.Writer,c.Request,nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()
		for{
		mt,message,error:=	ws.ReadMessage()
		if error!=nil {
			fmt.Println(error)
			break
		}
		if string(message) == "ping"{
			message = []byte("pong pong")
		}

		err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println(err)
				break
			}	

		}

	})
}