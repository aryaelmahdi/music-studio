package data

import (
	"context"
	"fmt"
	"net/smtp"
	"project/features/reservations"
	"project/features/users"

	"firebase.google.com/go/db"
	"firebase.google.com/go/messaging"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentData struct {
	client *snap.Client
	db     *db.Client
	fcm    *messaging.Client
}

func NewPaymentData(client *snap.Client, database *db.Client, messageClient *messaging.Client) *PaymentData {
	return &PaymentData{
		client: client,
		db:     database,
		fcm:    messageClient,
	}
}

func (pd *PaymentData) GetReservationInfo(reservationID string) (*reservations.Reservation, error) {
	ref := pd.db.NewRef("reservations").Child(reservationID)
	var reservation reservations.Reservation
	if err := ref.Get(context.Background(), &reservation); err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (pd *PaymentData) GetUserEmail(username string) (string, error) {
	ref := pd.db.NewRef("users").Child(username)
	var user users.User
	if err := ref.Get(context.Background(), &user); err != nil {
		return "", err
	}
	return user.Email, nil
}

func (pd *PaymentData) CreatePayment(chargeReq *snap.Request) (*snap.Response, error) {
	res, err := pd.client.CreateTransaction(chargeReq)
	if err != nil {
		fmt.Println("Error:", err.GetMessage())
		return nil, err
	}
	fmt.Println("Snap response:", res)
	return res, nil
}

func (pd *PaymentData) SendEmail(smtpUser string, smtpPassword string, smtpServer string, smtpPort string, receiver []string, msg string) error {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpUser, receiver, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (pd *PaymentData) SendMessage(message *messaging.Message) error {
	_, err := pd.fcm.Send(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
