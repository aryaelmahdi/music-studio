package handler

import (
	"net/http"
	"project/features/users"
	"project/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s users.UserServiceInterface
}

func NewUserHandler(service users.UserServiceInterface) users.UserHandlerInterface {
	return &UserHandler{
		s: service,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := users.User{}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}

		var user = new(users.User)
		user.Email = input.Email
		user.Username = input.Username
		user.Password = input.Password
		user.Role = "admin"

		if ok := strings.Contains(user.Email, "@"); !ok || len(user.Email) < 15 {
			c.Logger().Error("hanlder : invalid email format")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, invalid email format", nil, http.StatusBadRequest))
		}

		err := uh.s.Register(*user)

		var res = new(RegisterResponse)
		res.Username = user.Username

		if err != nil {
			c.Logger().Error("handler: input process error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res, http.StatusCreated))
	}
}

func (uh *UserHandler) AdminRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := users.User{}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}

		var user = new(users.User)
		user.Email = input.Email
		user.Username = input.Username
		user.Password = input.Password
		user.Role = "admin"

		if ok := strings.Contains(user.Email, "@"); !ok || len(user.Email) < 15 {
			c.Logger().Error("hanlder : invalid email format")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, invalid email format", nil, http.StatusBadRequest))
		}

		err := uh.s.Register(*user)

		var res = new(RegisterResponse)
		res.Username = user.Username

		if err != nil {
			c.Logger().Error("handler: input process error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res, http.StatusCreated))
	}
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		result, err := uh.s.Login(input.Username, input.Password)

		if err != nil {
			c.Logger().Error("handler: login process error:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("fail", nil, http.StatusNotFound))
			}
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}

		var res = new(LoginResponse)
		res.Username = result.Nama
		res.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}
