package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	//"fmt"
	"github.com/peterP1998/music-playlist-builder/service"
	"github.com/peterP1998/music-playlist-builder/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginController(t *testing.T) {
	userRepository:=mocks.UserRepositoryMock{}
	userService:=service.UserServiceCreate(userRepository)
	loginService:=service.LoginServiceAuth()
	loginController:=LoginControllerCreate(loginService,userService)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/login", loginController.Login)
	router.POST("/register", loginController.Register)
	t.Run("login with wrong credentials", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"user":"test","password":"test1"}`)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 401, resp.Result().StatusCode)
	})
	t.Run("login with correct credentials", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"user":"test1","password":"test"}`)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 200, resp.Result().StatusCode)
	})
	t.Run("register with existing user", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"user":"test","password":"test1","email":"test@abv.bg"}`)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 400, resp.Result().StatusCode)
	})
	t.Run("register with non existing user", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"user":"test21","password":"test1","email":"test@abv.bg"}`)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 200, resp.Result().StatusCode)
	})
}
