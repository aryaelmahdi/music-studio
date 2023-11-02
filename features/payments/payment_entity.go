package payments

import (
	"project/features/reservations"

	"firebase.google.com/go/messaging"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
)

type Payment struct {
	OrderID       string `json:"order_id"`
	ReservationID string `json:"reservation_id"`
	Username      string `json:"username"`
	GrossAmount   int    `json:"gross_amount"`
}

type PaymentData interface {
	CreatePayment(request *snap.Request) (*snap.Response, error)
	GetReservationInfo(reservationID string) (*reservations.Reservation, error)
	GetUserEmail(username string) (string, error)
	SendMessage(message *messaging.Message) error
	SendEmail(smtpUser string, smtpPassword string, smtpServer string, smtpPort string, receiver []string, msg string) error
	ConfirmedPayment(paymentInfo *Payment) error
	IsReservationValid(reservationID string) (bool, string)
	GetGrossAmount(roomID string) int
	UpdateStatus(status map[string]any, reservationID string) error
}
type PaymentService interface {
	CreatePayment(reservationID string) (*snap.Response, string, string, error)
	SendMessage(token string, paymentToken string, orderID string) error
	SendEmail(recipientEmail string, orderID string, paymentToken string) error
	ConfirmedPayment(orderID string) error
}

type PaymentHandler interface {
	CreatePayment() echo.HandlerFunc
	GetNotification() echo.HandlerFunc
}
