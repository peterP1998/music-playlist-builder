package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/controller"
	"net/http"
)

func main() {
	server := gin.New()
	var loginController controller.LoginController=controller.LoginController{}
	server.POST("/login", func(ctx *gin.Context) {
		token:= loginController.Login(ctx)
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	port := "8080"
	server.Run(":" + port)

}