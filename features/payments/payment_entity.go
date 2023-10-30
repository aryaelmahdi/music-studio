package payments

import (
	"project/features/reservations"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
)

type Payment struct {
	ReservationID string `json:"reservation_id"`
	Name          string `json:"name"`
	GrossAmount   int    `json:"gross_amount"`
}

type PaymentData interface {
	CreatePayment(reservationID string, username string, email string, price int) (*snap.Response, error)
	GetReservationInfo(reservationID string) (*reservations.Reservation, error)
	GetUserEmail(username string) (string, error)
}

type PaymentService interface {
	CreatePayment(reservationID string) (*snap.Response, error)
}

type PaymentHandler interface {
	CreatePayment() echo.HandlerFunc
}
