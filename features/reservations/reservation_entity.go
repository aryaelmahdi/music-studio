package reservations

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Reservation struct {
	ReservationID string `json:"reservation_id"`
	RoomID        string `json:"room_id"`
	Username      string `json:"username"`
	Price         int    `json:"price"`
	PaymentStatus string `json:"payment_status"`
}

type AllReservations map[string]interface{}

type ReservationData interface {
	GetAllReservations() (*AllReservations, error)
	GetReservationsByUsername(username string) (map[string]any, error)
	AddReservation(newData Reservation) (*Reservation, error)
}

type ReservationService interface {
	GetAllReservations(token *jwt.Token) (*AllReservations, error)
	GetReservationsByUsername(token *jwt.Token) (map[string]any, error)
	AddReservation(newData Reservation) (*Reservation, error)
}

type ReservationHandler interface {
	GetAllReservations() echo.HandlerFunc
	GetReservationsByUsername() echo.HandlerFunc
	AddReservation() echo.HandlerFunc
}
