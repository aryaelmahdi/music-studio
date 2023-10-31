package service

import (
	"errors"
	"project/features/payments"
	"project/helper"

	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService struct {
	d payments.PaymentData
	j helper.JWTInterface
}

func NewPaymentService(data payments.PaymentData, jwt helper.JWTInterface) *PaymentService {
	return &PaymentService{
		d: data,
		j: jwt,
	}
}

func (ps *PaymentService) CreatePayment(reservationID string) (*snap.Response, error) {
	reservation, err := ps.d.GetReservationInfo(reservationID)
	if err != nil {
		return nil, err
	}

	if reservation.Username == "" {
		return nil, errors.New("No data")
	}

	email, err := ps.d.GetUserEmail(reservation.Username)
	if err != nil {
		return nil, err
	}

	res, err := ps.d.CreatePayment(reservationID, reservation.Username, email, int(reservation.Price.(float64)))
	if err != nil {
		return nil, err
	}
	return res, nil
}
