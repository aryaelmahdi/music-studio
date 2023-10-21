package routes

import (
	"project/features/users"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc users.UserHandlerInterface) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
}
