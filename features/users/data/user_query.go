package data

import (
	"context"
	"errors"
	"project/features/users"

	"firebase.google.com/go/db"
	"github.com/sirupsen/logrus"
)

type UserData struct {
	db *db.Client
}

func NewUserData(client *db.Client) users.UserDataInterface {
	return &UserData{
		db: client,
	}
}

func (ud *UserData) Insert(newData users.User) error {
	userError := errors.New("user exists")
	exist := ud.GetUserByUsername(newData.Username)
	if !exist {
		return userError
	}
	ref := ud.db.NewRef("users").Child(newData.Username)
	if err := ref.Set(context.Background(), newData); err != nil {
		return err
	}

	return nil
}

func (ud *UserData) Login(username string, password string) (*users.User, error) {
	ref := ud.db.NewRef("users").Child(username)
	var user users.User
	if err := ref.Get(context.Background(), &user); err != nil {
		return nil, err
	}
	if user.Password != password {
		logrus.Info("users :", user)
		logrus.Info("users password :", user.Password)
		return nil, nil
	}

	return &user, nil
}

func (ud *UserData) GetUserByUsername(username string) bool {
	ref := ud.db.NewRef("users").Child(username)
	var user users.User
	if err := ref.Get(context.Background(), &user); err != nil {
		return false
	}
	return true
}
