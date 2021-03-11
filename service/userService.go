package service

import (
	"fmt"
	"github.com/peterP1998/music-playlist-builder/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	SelectUserByName(username string) (model.User, error)
	CreateUser(username string, password string, email string) error
}

type UserServiceInterface interface {
	AuthenticateUser(username string, password string) (bool, error)
	RegisterUser(username string, email string, password string) (bool, error)
}

type UserService struct {
	userRepository UserRepositoryInterface
}

func UserServiceCreate(userRepository UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
	}
}
func (userService *UserService) AuthenticateUser(username string, password string) (bool, error) {
	user, err := userService.userRepository.SelectUserByName(username)
	if err != nil {
		return false, err
	}
	if user.Username == username && (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))) == nil {
		return true, nil
	}
	return false, nil
}

func (userService *UserService) RegisterUser(username string, email string, password string) (bool, error) {
	exist, err := checkDoesUserAlreadyExists(userService, username)
	if err != nil || exist {
		return false, err
	}
	hashPassword, err := hashPassword(password)
	if err != nil {
		return false, err
	}
	err = userService.userRepository.CreateUser(username, email, hashPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}

func checkDoesUserAlreadyExists(userService *UserService, username string) (bool, error) {
	_, err := userService.userRepository.SelectUserByName(username)
	if err != nil || fmt.Sprint(err) == "sql: no rows in result set" {
		return false, nil
	}
	return true, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}
