package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/model/requests"
	"github.com/peterP1998/music-playlist-builder/service"
	"net/http"
	"fmt"
)

type LoginControllerInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
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
	var credential requests.UserCrednetials
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

func (controller *LoginController) Register(ctx *gin.Context) {
	var userRequest requests.UserRegister
	err:= ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	register,err:= controller.userService.RegisterUser(userRequest.Username,userRequest.Email ,userRequest.Password)
	if !register || err!=nil  {
		fmt.Println(err)
		ctx.JSON(500, "Something went wrong!")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"info": "Registration is successful!",
	})
}