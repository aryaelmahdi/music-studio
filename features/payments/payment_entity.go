package payments

import (
	"project/features/reservations"

	"firebase.google.com/go/messaging"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
)

type Payment struct {
	ReservationID string `json:"reservation_id"`
	Name          string `json:"name"`
	GrossAmount   int    `json:"gross_amount"`
}

type PaymentData interface {
	CreatePayment(request *snap.Request) (*snap.Response, error)
	GetReservationInfo(reservationID string) (*reservations.Reservation, error)
	GetUserEmail(username string) (string, error)
	SendMessage(message *messaging.Message) error
	SendEmail(smtpUser string, smtpPassword string, smtpServer string, smtpPort string, receiver []string, msg string) error
}
type PaymentService interface {
	CreatePayment(reservationID string) (*snap.Response, string, string, error)
	SendMessage(token string, paymentToken string, orderID string) error
	SendEmail(recipientEmail string, orderID string, paymentToken string) error
}

type PaymentHandler interface {
	CreatePayment() echo.HandlerFunc
}
