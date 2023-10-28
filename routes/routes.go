package routes

import (
	"project/features/instruments"
	"project/features/payments"
	"project/features/reservations"
	"project/features/rooms"
	"project/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
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

func InstrumentsRoutes(e *echo.Echo, ih instruments.InstrumentHandler, secret string) {
	e.POST("/instruments", ih.AddInstrument(), echojwt.JWT([]byte(secret)))
	e.GET("/instruments", ih.GetAllInstruments())
	e.GET("/instruments/:id", ih.GetInstrumentByID())
	e.PUT("/instruments/:id", ih.UpdateInstrument(), echojwt.JWT([]byte(secret)))
	e.DELETE("instruments/:id", ih.DeleteInstrument(), echojwt.JWT([]byte(secret)))
}

func ReservationRoutes(e *echo.Echo, rh reservations.ReservationHandler, secret string) {
	e.POST("/reservations", rh.AddReservation(), echojwt.JWT([]byte(secret)))
	e.GET("/reservations", rh.GetAllReservations(), echojwt.JWT([]byte(secret)))
	e.GET("/myreservations", rh.GetReservationsByUsername(), echojwt.JWT([]byte(secret)))
}

func PaymentRoutes(e *echo.Echo, ph payments.PaymentHandler, secret string) {
	e.GET("/payment/:id", ph.CreatePayment())
}
