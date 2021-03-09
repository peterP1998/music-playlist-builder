package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/model"
	"github.com/peterP1998/music-playlist-builder/service"
	"net/http"
)

type LoginControllerInterface interface {
	Login(ctx *gin.Context)
}

type LoginController struct {
	userService service.UserServiceInterface
	loginService service.LoginServiceInterface
}

func LoginControllerCreate(loginService service.LoginServiceInterface,userService service.UserServiceInterface) LoginControllerInterface {
	return &LoginController{
		loginService: loginService,
		userService:    userService,
	}
}

func (controller *LoginController) Login(ctx *gin.Context) {
	var credential model.UserCrednetials
	err:= ctx.ShouldBindJSON(&credential)
	if err != nil {
		ctx.JSON(400, nil)
		return
	}
	isUserAuthenticated,_:= controller.userService.AuthenticateUser(credential.Username, credential.Password)
	token:=""
	if isUserAuthenticated {
		token=controller.loginService.GenerateToken(credential.Username, true)
	}
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}
