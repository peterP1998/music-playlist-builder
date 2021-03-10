package repository

import (
	"github.com/peterP1998/music-playlist-builder/model"
)

type UserRepository struct {
}


func (ur UserRepository) SelectUserByName(username string) (model.User, error) {
	var user model.User
	err := model.DB.QueryRow("SELECT * FROM User where username=?", username).Scan(&user.Id, &user.Email, &user.Username,  &user.Password)
	if err!=nil {
		return user, err
	}
	return user, nil
}

func (ur UserRepository) CreateUser(username string,email string,password string) error {
	_, err := model.DB.Query("insert into User(email,username,password) Values(?,?,?);", email,username, password)
	if err != nil {
		return err
	}
	return nil
}