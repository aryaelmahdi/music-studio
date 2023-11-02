package service

import (
	"errors"
	"fmt"
	"project/config"
	"project/features/payments"
	"project/helper"

	"firebase.google.com/go/messaging"
	"github.com/midtrans/midtrans-go"
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

func (ps *PaymentService) CreatePayment(reservationID string) (*snap.Response, string, string, error) {
	reservation, err := ps.d.GetReservationInfo(reservationID)
	if err != nil {
		return nil, "", "", err
	}

	if reservation.Username == "" {
		return nil, "", "", errors.New("No data")
	}

	email, err := ps.d.GetUserEmail(reservation.Username)
	if err != nil {
		return nil, "", "", err
	}
	snapRequest := generateSnapReq(reservationID, reservation.Username, email, int(reservation.Price.(float64)))
	res, err := ps.d.CreatePayment(snapRequest)
	if err != nil {
		return nil, "", "", err
	}
	return res, snapRequest.TransactionDetails.OrderID, email, nil
}

func (ps *PaymentService) ConfirmedPayment(orderID string) error {
	reservationID := orderID[3:]
	reservationValid, username := ps.d.IsReservationValid(reservationID)
	if !reservationValid || username == "" {
		return errors.New("Reservation not found")
	}
	roomID := orderID[12:]
	grossAmount := ps.d.GetGrossAmount(roomID)

	var paymentInfo payments.Payment
	paymentInfo.GrossAmount = grossAmount
	paymentInfo.OrderID = orderID
	paymentInfo.ReservationID = reservationID
	paymentInfo.Username = username

	if err := ps.d.ConfirmedPayment(&paymentInfo); err != nil {
		return err
	}

	status := make(map[string]any)
	status["payment_status"] = "finished"

	if err := ps.d.UpdateStatus(status, reservationID); err != nil {
		return err
	}

	return nil
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
		"your Order ID : " + orderID + "\n" +
		"Please complete your payment..."

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

func (ps *PaymentService) ConfirmedPaymentEmail(orderID string) error {
	reservationID := orderID[3:]
	reservationValid, username := ps.d.IsReservationValid(reservationID)
	if !reservationValid || username == "" {
		return errors.New("Reservation not found")
	}

	email, err := ps.d.GetUserEmail(username)
	if err != nil {
		return errors.New("user not found")
	}
	to := make([]string, 1)
	to[0] = email
	subject := "Payment Completed"
	body := "Hello Users!! \n \n" +
		"Thank You for completing the payment, \n" +
		"your Order ID : " + orderID + "\n \n" +
		"Enjoy Musicians!"

	message := "From: " + ps.c.From + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	fmt.Println(ps.c.EmailUsername, ps.c.EmailPassword, ps.c.EmailHost, ps.c.EmailPort, to, message)
	sendErr := ps.d.SendEmail(ps.c.EmailUsername, ps.c.EmailPassword, ps.c.EmailHost, ps.c.EmailPort, to, message)
	if sendErr != nil {
		return sendErr
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

func generateSnapReq(reservationID string, username string, email string, price int) *snap.Request {
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "GSM" + reservationID,
			GrossAmt: int64(price),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: username,
			Email: email,
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}
	return snapReq
}
