package data

import (
	"context"
	"errors"
	"fmt"
	"net/smtp"
	"project/features/payments"
	"project/features/reservations"
	"project/features/rooms"
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

func (pd *PaymentData) ConfirmedPayment(paymentInfo *payments.Payment) error {
	ref := pd.db.NewRef("payments").Child(paymentInfo.OrderID)
	if err := ref.Set(context.Background(), &paymentInfo); err != nil {
		return errors.New("something went wrong")
	}
	return nil
}

func (pd *PaymentData) IsReservationValid(reservationID string) (bool, string) {
	ref := pd.db.NewRef("reservations").Child(reservationID)
	var reservation reservations.Reservation
	ref.Get(context.Background(), &reservation)
	if reservation.Username == "" {
		return false, ""
	}
	return true, reservation.Username
}

func (pd *PaymentData) GetGrossAmount(roomID string) int {
	ref := pd.db.NewRef("rooms").Child(roomID)
	var room rooms.Rooms
	ref.Get(context.Background(), &room)
	return room.Price
}

func (pd *PaymentData) UpdateStatus(status map[string]any, reservationID string) error {
	ref := pd.db.NewRef("reservations").Child(reservationID)
	if err := ref.Update(context.Background(), status); err != nil {
		return errors.New("cannot update payment status")
	}
	return nil
}
