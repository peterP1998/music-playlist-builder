package service

import (
	"github.com/peterP1998/music-playlist-builder/model"
)

type UserRepositoryInterface interface {
	SelectUserByName(username string) (model.User, error)
}

type UserServiceInterface interface {
	AuthenticateUser(username string, password string) (bool, error)
}

type UserService struct {
	userRepository UserRepositoryInterface
}
func UserServiceCreate(userRepository UserRepositoryInterface) UserServiceInterface{
	return &UserService{
		userRepository: userRepository,
	}
}
func (userService *UserService) AuthenticateUser(username string, password string) (bool, error) {
	user, err := userService.userRepository.SelectUserByName(username)
	if err != nil {
		return false, err
	}
	if user.Username == username &&user.Password==password {
		return true, nil
	}
	return false, nil
}
