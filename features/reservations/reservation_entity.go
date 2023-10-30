package reservations

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Reservation struct {
	ReservationID string `json:"reservation_id"`
	RoomID        string `json:"room_id"`
	Username      string `json:"username"`
	Price         any    `json:"price"`
	PaymentStatus string `json:"payment_status"`
	Date          string `json:"date"`
}

type AllReservations map[string]interface{}

type ReservationData interface {
	GetAllReservations() (*AllReservations, error)
	GetReservationsByUsername(username string) (map[string]any, error)
	AddReservation(newData Reservation) (*Reservation, error)
	UpdateReservation(newData map[string]interface{}) (*Reservation, error)
	DeleteReservation(id string) error
	GetReservationByID(id string) (*Reservation, error)
}

type ReservationService interface {
	GetAllReservations(token *jwt.Token) (*AllReservations, error)
	GetReservationsByUsername(token *jwt.Token) (map[string]any, error)
	AddReservation(newData Reservation, token *jwt.Token) (*Reservation, error)
	UpdateReservation(newData Reservation, token *jwt.Token) (*Reservation, error)
	DeleteReservation(id string, token *jwt.Token) error
	GetReservationByID(id string) (*Reservation, error)
}

type ReservationHandler interface {
	GetAllReservations() echo.HandlerFunc
	GetReservationsByUsername() echo.HandlerFunc
	AddReservation() echo.HandlerFunc
	UpdateReservation() echo.HandlerFunc
	DeleteReservation() echo.HandlerFunc
}
