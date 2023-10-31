package service

import (
	"errors"
	"fmt"
	"project/config"
	"project/features/payments"
	"project/helper"

	"firebase.google.com/go/messaging"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService struct {
	d payments.PaymentData
	j helper.JWTInterface
	c config.SMTP
}

func NewPaymentService(data payments.PaymentData, jwt helper.JWTInterface, cfg config.SMTP) *PaymentService {
	return &PaymentService{
		d: data,
		j: jwt,
		c: cfg,
	}
}

func (ps *PaymentService) CreatePayment(reservationID string) (*snap.Response, string, error) {
	reservation, err := ps.d.GetReservationInfo(reservationID)
	if err != nil {
		return nil, "", err
	}

	if reservation.Username == "" {
		return nil, "", errors.New("No data")
	}

	email, err := ps.d.GetUserEmail(reservation.Username)
	if err != nil {
		return nil, "", err
	}

	res, orderID, err := ps.d.CreatePayment(reservationID, reservation.Username, email, int(reservation.Price.(float64)))
	if err != nil {
		return nil, "", err
	}
	return res, orderID, nil
}

func (ps *PaymentService) SendMessage(token string, paymentToken string, orderID string) error {
	message := generateMessage(token, paymentToken, paymentToken)
	if err := ps.d.SendMessage(message); err != nil {
		return err
	}
	return nil
}

func (ps *PaymentService) SendEmail(recipientEmail string, orderID string, paymentToken string) error {
	to := make([]string, 1)
	to[0] = recipientEmail
	subject := "Payment Details"
	body := "Hello Users!! \n" +
		"Thank You for reserving studio, \n" +
		"your Order ID = " + orderID + "\n" +
		"here's your payment token : \n" + paymentToken + "\n" +
		"Please confirm your payment using this token"

	message := "From: " + ps.c.From + "\n" +
		"To: " + recipientEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	fmt.Println(ps.c.EmailUsername, ps.c.EmailPassword, ps.c.EmailHost, ps.c.EmailPort, to, message)
	err := ps.d.SendEmail(ps.c.EmailUsername, ps.c.EmailPassword, ps.c.EmailHost, ps.c.EmailPort, to, message)
	if err != nil {
		return err
	}
	return nil
}

func generateMessage(token string, paymentToken string, orderID string) *messaging.Message {
	notification := &messaging.Notification{
		Title: "Payment Notification ",
		Body: "Hello Users!! \n" +
			"Thank You for reserving studio, \n" +
			"your Order ID = " + orderID + "\n" +
			"here's your payment token : \n" + paymentToken +
			"Please confirm your payment using this token",
	}

	message := &messaging.Message{
		Notification: notification,
		Token:        token,
	}
	return message
}
