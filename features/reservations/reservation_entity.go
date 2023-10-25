package reservations

import "github.com/labstack/echo/v4"

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
	GetReservationsByUsername(username string) (*Reservation, error)
	AddReservation(newData Reservation) (*Reservation, error)
}

type ReservationService interface {
	GetAllReservations() (*AllReservations, error)
	GetReservationsByUsername(username string) (*Reservation, error)
	AddReservation(newData Reservation) (*Reservation, error)
}

type ReservationHandler interface {
	GetAllReservations() echo.HandlerFunc
	GetReservationsByUsername() echo.HandlerFunc
	AddReservation() echo.HandlerFunc
}
