package controllers

import (
	"github.com/gin-gonic/gin"
)



type ChatController struct{
}


func NewChatController( ) *ChatController{
	return &ChatController{
	
	}
}

func (cC * ChatController) Pong(request *gin.Context ){
	request.JSON(200,gin.H{
		"message":"pong",
	})
}

func (cC *ChatController) HandleChat(request *gin.Context){




}