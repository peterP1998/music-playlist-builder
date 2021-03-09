package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/controller"
	"github.com/peterP1998/music-playlist-builder/service"
	"github.com/peterP1998/music-playlist-builder/repository"
	"github.com/peterP1998/music-playlist-builder/model"
	"database/sql"
)

func main() {
	model.DB,_=sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/music")
	server := gin.New()
	var userRepository repository.UserRepository=repository.UserRepository{}
	var userService service.UserServiceInterface=service.UserServiceCreate(userRepository)
	var loginService service.LoginServiceInterface=service.LoginServiceAuth()
	var loginController controller.LoginControllerInterface=controller.LoginControllerCreate(loginService,userService)
	server.POST("/login", loginController.Login)
	port := "8080"
	server.Run(":" + port)

}