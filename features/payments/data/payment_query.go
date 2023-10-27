package data

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentData struct {
	client *snap.Client
}

func NewPaymentData(client snap.Client) *PaymentData {
	return &PaymentData{
		client: &client,
	}
}

func (pd *PaymentData) CreatePayment(id string) (*snap.Response, error) {
	chargeReq := generateSnapReq(id)
	res, err := pd.client.CreateTransaction(chargeReq)
	if err != nil {
		fmt.Println("Error:", err.GetMessage())
		return nil, err
	}
	fmt.Println("Snap response:", res)
	return res, nil
}

func generateSnapReq(reservationID string) *snap.Request {
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  reservationID,
			GrossAmt: 1,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "arya",
			Email: "elprab17@gmail.com",
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}
	return snapReq
}
