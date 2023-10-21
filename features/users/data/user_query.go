package data

import (
	"context"
	"fmt"
	"project/features/users"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/sirupsen/logrus"
)

type UserData struct {
	db  *db.Client
	app *firebase.App
}

func NewUserData(client *db.Client, app *firebase.App) users.UserDataInterface {
	return &UserData{
		db:  client,
		app: app,
	}
}

func (ud *UserData) Insert(newData users.User) error {
	fmt.Println(newData)
	ref := ud.db.NewRef("user").Child(newData.Username)
	if err := ref.Set(context.Background(), newData); err != nil {
		return err
	}

	return nil
}

func (ud *UserData) Login(username string, password string) (*users.User, error) {
	ref := ud.db.NewRef("user/").Child(username)
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
