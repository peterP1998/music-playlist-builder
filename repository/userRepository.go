package repository

import (
	"github.com/peterP1998/music-playlist-builder/model"
)

type UserRepository struct {
}


func (ur UserRepository) SelectUserByName(username string) (model.User, error) {
	var user model.User
	err := model.DB.QueryRow("SELECT * FROM User where username=?", username).Scan(&user.Id, &user.Email, &user.Username,  &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
