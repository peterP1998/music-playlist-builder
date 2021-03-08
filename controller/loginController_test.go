package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	resp := httptest.NewRecorder()
	var loginController LoginController=LoginController{}
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)
	token:=loginController.Login(c)
	assert.Equal(t, "", token)
}
