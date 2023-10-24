package routes

import (
	"project/features/rooms"
	"project/features/users"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uh users.UserHandlerInterface) {
	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())
}

func RoomRoutes(e *echo.Echo, rh rooms.RoomHandler, secret string) {
	e.POST("/rooms", rh.AddRoom(), echojwt.JWT([]byte(secret)))
	e.GET("/rooms", rh.GetAllRooms())
	e.GET("/rooms/:id", rh.GetRoomByID())
	e.DELETE("/rooms/:id", rh.DeleteRoom(), echojwt.JWT([]byte(secret)))
	e.PUT("/rooms/:id", rh.UpdateRoom(), echojwt.JWT([]byte(secret)))
}
