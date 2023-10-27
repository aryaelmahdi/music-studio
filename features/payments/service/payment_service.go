package service

import (
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

func (ps *PaymentService) CreatePayment(id string) (*snap.Response, error) {
	res, err := ps.d.CreatePayment(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
