package service

import (
	"errors"
	"project/features/users"
	"project/features/users/mocks"
	helperMock "project/helper/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsert(t *testing.T) {
	jwt := helperMock.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	var service = NewUserService(data, jwt)

	var newData = users.User{
		Username: "arya",
		Password: "gugu",
		Email:    "arya@gmail.com",
	}

	invalidData := users.User{
		Username: "arya",
		Password: "",
	}

	t.Run("success insert", func(t *testing.T) {
		data.On("Insert", newData).Return(nil).Once()
		err := service.Register(newData)
		assert.Nil(t, err)
	})

	t.Run("Failed insert", func(t *testing.T) {
		data.On("Insert", invalidData).Return(errors.New("insert process failed")).Once()
		err := service.Register(invalidData)
		assert.Error(t, err)
		assert.NotNil(t, err)
	})

	data.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	jwt := helperMock.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	var service = NewUserService(data, jwt)

	var newData = users.User{
		Username: "arya",
		Password: "gugu",
		Email:    "arya@gmail.com",
		Role:     "admin",
	}

	invalidData := users.User{
		Username: "arya",
		Password: "",
	}

	t.Run("success login", func(t *testing.T) {
		var jwtResult = map[string]any{"access_token": "randomAccessToken"}
		data.On("Login", newData.Username, newData.Password).Return(&newData, nil).Once()
		jwt.On("GenerateJWT", mock.Anything, mock.Anything).Return(jwtResult).Once()

		result, err := service.Login(newData.Username, newData.Password)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "arya", result.Nama)
		assert.Equal(t, jwtResult, result.Access)
		data.AssertExpectations(t)
		jwt.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		data.On("Login", invalidData.Username, invalidData.Password).Return(nil, errors.New("password does not matched!")).Once()

		res, err := service.Login(invalidData.Username, invalidData.Password)
		assert.NotNil(t, err)
		assert.Nil(t, res)
		data.AssertExpectations(t)
	})

	data.AssertExpectations(t)
}
