package service

import (
	"errors"
	"fmt"
	"project/features/users"
	"project/helper"
	"strings"
)

type UserService struct {
	d users.UserDataInterface
	g helper.GeneratorInterface
	j helper.JWTInterface
}

func NewUserService(data users.UserDataInterface, generator helper.GeneratorInterface, jwt helper.JWTInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		g: generator,
		j: jwt,
	}
}

func (us *UserService) Register(newData users.User) error {
	err := us.d.Insert(newData)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Login(username string, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(username, password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("process failed")
	}
	if result == nil && err == nil {
		return nil, errors.New("password does not matched!")
	}
	fmt.Println("service :", result)

	tokenData := us.j.GenerateJWT(result.Username, result.Role)

	if tokenData == nil {
		return nil, errors.New("token process failed")
	}

	response := new(users.UserCredential)
	response.Nama = result.Username
	response.Access = tokenData

	return response, nil
}
