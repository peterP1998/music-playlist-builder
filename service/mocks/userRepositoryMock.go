package mocks

import (
	"errors"
	"github.com/peterP1998/music-playlist-builder/model"
)

type UserRepositoryMock struct {
}

func (ur UserRepositoryMock) SelectUserByName(username string) (model.User, error) {
	var user model.User
	if username == "test" {
		return getUser(1, "test", "test", "test@test.bg"), nil
	} else if username == "test1" {
		return getUser(2, "test1", "test1", "test1@test1.bg"), nil
	} else if username == "test2" {
		return user, errors.New("User not found")
	} else {
		return model.User{}, nil
	}
}

func (ur UserRepositoryMock) CreateUser(username string, password string, email string) error {
	if username == "test21" {
		return nil
	}
	return errors.New("User is not created")
}

func getUser(id int64, username string, pass string, email string) model.User {
	var user model.User
	user.Id = id
	user.Username = username
	user.Email = email
	user.Password = pass
	return user
}
