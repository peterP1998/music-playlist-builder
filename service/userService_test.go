package service

import (
	"github.com/stretchr/testify/assert"
	//"net/http/httptest"
	"github.com/peterP1998/music-playlist-builder/service/mocks"
	"testing"
)

func TestAuthenticateUserWithExistingUser(t *testing.T) {
	userRepositoryMock := mocks.UserRepositoryMock{}
	userService := UserServiceCreate(userRepositoryMock)
	auth, err := userService.AuthenticateUser("test1", "test")
	assert.Equal(t, true, auth)
	assert.Equal(t, nil, err)
}

func TestAuthenticateUserWithNonExistingUser(t *testing.T) {
	userRepositoryMock := mocks.UserRepositoryMock{}
	userService := UserServiceCreate(userRepositoryMock)
	auth, err := userService.AuthenticateUser("test2", "test")
	assert.Equal(t, false, auth)
	assert.NotEqual(t, nil, err)
}

func TestAuthenticateUserWithExistingUserAndWrongCredentials(t *testing.T) {
	userRepositoryMock := mocks.UserRepositoryMock{}
	userService := UserServiceCreate(userRepositoryMock)
	auth, err := userService.AuthenticateUser("test", "test1")
	assert.Equal(t, false, auth)
	assert.Equal(t, nil, err)
}

func TestRegisterUserWithExistingUser(t *testing.T) {
	userRepositoryMock := mocks.UserRepositoryMock{}
	userService := UserServiceCreate(userRepositoryMock)
	register, err := userService.RegisterUser("test", "test1@test1.bg", "test1")
	assert.Equal(t, false, register)
	assert.NotEqual(t, nil, err)
}

func TestRegisterUserWithNonExistingUser(t *testing.T) {
	userRepositoryMock := mocks.UserRepositoryMock{}
	userService := UserServiceCreate(userRepositoryMock)
	register, err := userService.RegisterUser("test21", "test1@test1.bg", "test1")
	assert.Equal(t, true, register)
	assert.Equal(t, nil, err)
}

func TestGetUserExistingUser(t *testing.T) {
	userRepositoryMock := mocks.UserRepositoryMock{}
	userService := UserServiceCreate(userRepositoryMock)
	user, err := userService.GetUser("test")
	assert.Equal(t, "test", user.Username)
	assert.Equal(t, nil, err)
}

func TestGetUserNonExistingUser(t *testing.T) {
	userRepositoryMock := mocks.UserRepositoryMock{}
	userService := UserServiceCreate(userRepositoryMock)
	_, err := userService.GetUser("test2")
	assert.NotEqual(t, nil, err)
}
