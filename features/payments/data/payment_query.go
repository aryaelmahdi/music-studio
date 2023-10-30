package data

import (
	"context"
	"fmt"
	"project/features/reservations"
	"project/features/users"

	"firebase.google.com/go/db"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentData struct {
	client *snap.Client
	db     *db.Client
}

func NewPaymentData(client snap.Client, database *db.Client) *PaymentData {
	return &PaymentData{
		client: &client,
		db:     database,
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

func (pd *PaymentData) CreatePayment(reservationID string, username string, email string, price int) (*snap.Response, error) {
	chargeReq := generateSnapReq(reservationID, username, email, price)
	fmt.Println("price : ", price)
	res, err := pd.client.CreateTransaction(chargeReq)
	if err != nil {
		fmt.Println("Error:", err.GetMessage())
		return nil, err
	}
	fmt.Println("Snap response:", res)
	return res, nil
}

func generateSnapReq(reservationID string, username string, email string, price int) *snap.Request {
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  reservationID,
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
