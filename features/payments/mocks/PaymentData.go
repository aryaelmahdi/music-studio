// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	messaging "firebase.google.com/go/messaging"
	mock "github.com/stretchr/testify/mock"

	payments "project/features/payments"

	reservations "project/features/reservations"

	snap "github.com/midtrans/midtrans-go/snap"
)

// PaymentData is an autogenerated mock type for the PaymentData type
type PaymentData struct {
	mock.Mock
}

// ConfirmedPayment provides a mock function with given fields: paymentInfo
func (_m *PaymentData) ConfirmedPayment(paymentInfo *payments.Payment) error {
	ret := _m.Called(paymentInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(*payments.Payment) error); ok {
		r0 = rf(paymentInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePayment provides a mock function with given fields: request
func (_m *PaymentData) CreatePayment(request *snap.Request) (*snap.Response, error) {
	ret := _m.Called(request)

	var r0 *snap.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*snap.Request) (*snap.Response, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*snap.Request) *snap.Response); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snap.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*snap.Request) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGrossAmount provides a mock function with given fields: roomID
func (_m *PaymentData) GetGrossAmount(roomID string) int {
	ret := _m.Called(roomID)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(roomID)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetReservationInfo provides a mock function with given fields: reservationID
func (_m *PaymentData) GetReservationInfo(reservationID string) (*reservations.Reservation, error) {
	ret := _m.Called(reservationID)

	var r0 *reservations.Reservation
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*reservations.Reservation, error)); ok {
		return rf(reservationID)
	}
	if rf, ok := ret.Get(0).(func(string) *reservations.Reservation); ok {
		r0 = rf(reservationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*reservations.Reservation)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(reservationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserEmail provides a mock function with given fields: username
func (_m *PaymentData) GetUserEmail(username string) (string, error) {
	ret := _m.Called(username)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsReservationValid provides a mock function with given fields: reservationID
func (_m *PaymentData) IsReservationValid(reservationID string) (bool, string) {
	ret := _m.Called(reservationID)

	var r0 bool
	var r1 string
	if rf, ok := ret.Get(0).(func(string) (bool, string)); ok {
		return rf(reservationID)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(reservationID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(reservationID)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// SendEmail provides a mock function with given fields: smtpUser, smtpPassword, smtpServer, smtpPort, receiver, msg
func (_m *PaymentData) SendEmail(smtpUser string, smtpPassword string, smtpServer string, smtpPort string, receiver []string, msg string) error {
	ret := _m.Called(smtpUser, smtpPassword, smtpServer, smtpPort, receiver, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, []string, string) error); ok {
		r0 = rf(smtpUser, smtpPassword, smtpServer, smtpPort, receiver, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMessage provides a mock function with given fields: message
func (_m *PaymentData) SendMessage(message *messaging.Message) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(*messaging.Message) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: status, reservationID
func (_m *PaymentData) UpdateStatus(status map[string]interface{}, reservationID string) error {
	ret := _m.Called(status, reservationID)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}, string) error); ok {
		r0 = rf(status, reservationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPaymentData creates a new instance of PaymentData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentData(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentData {
	mock := &PaymentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
