package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/model"
)

type LoginControllerInterface interface {
	Login(ctx *gin.Context) string
}

type LoginController struct {
}

func (controller *LoginController) Login(ctx *gin.Context) string {
	var credential model.UserCrednetials
	_ = ctx.ShouldBindJSON(&credential)
	fmt.Println(credential.Username)
	return credential.Username
}
