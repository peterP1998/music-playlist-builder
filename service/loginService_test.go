package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	loginService := LoginServiceAuth()
	token, err := loginService.GenerateToken("test@tets.bg")
	assert.NotEqual(t, "", token)
	assert.Equal(t, nil, err)
}

func TestValidateTokenWithFakeToken(t *testing.T) {
	loginService := LoginServiceAuth()
	_, err := loginService.ValidateToken("test")
	assert.NotEqual(t, "Invalid token", fmt.Sprint(err))
}

func TestValidateTokenWithRealToken(t *testing.T) {
	loginService := LoginServiceAuth()
	token, err := loginService.GenerateToken("test@tets.bg")
	_, err = loginService.ValidateToken(token)
	assert.NotEqual(t, "Invalid token", fmt.Sprint(err))
}
