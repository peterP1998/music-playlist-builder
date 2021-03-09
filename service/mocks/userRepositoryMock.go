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
		return getUser(1, "test", "test@test.bg", "test"), nil
	} else if username == "test1" {
		return getUser(2, "test1", "test@test.bg", "test"), nil
	} else {
		return user, errors.New("User not found")
	}
}
func getUser(id int64, username string, pass string, email string) model.User {
	var user model.User
	user.Id = id
	user.Username = username
	user.Email = email
	user.Password = pass
	return user
}
