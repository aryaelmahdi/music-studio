package users

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserCredential struct {
	Nama   string
	Access map[string]any
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	AdminRegister() echo.HandlerFunc
}
type UserServiceInterface interface {
	Register(newData User) error
	Login(username string, password string) (*UserCredential, error)
}

type UserDataInterface interface {
	Insert(newData User) error
	Login(username string, password string) (*User, error)
}
