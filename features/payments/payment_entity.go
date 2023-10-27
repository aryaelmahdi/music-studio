package payments

import (
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
)

type Payment struct {
	OrderID     string `json:"order_id"`
	Name        string `json:"name"`
	GrossAmount int    `json:"gross_amount"`
}

type PaymentData interface {
	CreatePayment(id string) (*snap.Response, error)
}

type PaymentService interface {
	CreatePayment(id string) (*snap.Response, error)
}

type PaymentHandler interface {
	CreatePayment() echo.HandlerFunc
}
