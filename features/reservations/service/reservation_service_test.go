package service

import (
	"project/features/reservations"
	"project/features/reservations/mocks"
	helperMock "project/helper/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func NewHelperDate(t *testing.T) *helperMock.HelperDate {
	mock := &helperMock.HelperDate{}
	mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func TestReservationService_GetAllReservation(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewReservationData(t)
	service := NewReservationService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := reservations.Reservation{}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		jwtt.On("ExtractToken", mockToken).Return("id", "role")
		data.On("GetAllReservations").Return(newData, nil)

		_, err := service.GetAllReservations(mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}
func TestReservationService_GetReservationByUsername(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewReservationData(t)
	service := NewReservationService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := reservations.Reservation{}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		claim := mockToken.Claims.(jwt.MapClaims)
		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("GetReservationsByUsername", mockToken).Return(newData, nil)

		_, err := service.GetReservationsByUsername(mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}
func TestReservationService_AddReservation(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewReservationData(t)
	service := NewReservationService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := reservations.Reservation{}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		claim := mockToken.Claims.(jwt.MapClaims)
		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("AddReservation", mockToken).Return(newData, nil)

		_, err := service.AddReservation(newData, mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestReservationService_UpdateReservation(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewReservationData(t)
	date := NewHelperDate(t)
	service := NewReservationService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := reservations.Reservation{}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		id := "1"
		claim := mockToken.Claims.(jwt.MapClaims)
		resDate := "20-12-2024"
		date.On("CompareDate", resDate).Return(nil)
		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("UpdateReservation", mockToken).Return(newData, nil)

		_, err := service.UpdateReservation(id, newData, mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestReservationService_DeleteReservation(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewReservationData(t)
	service := NewReservationService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		id := "1"

		claim := mockToken.Claims.(jwt.MapClaims)

		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("DeleteReservation", mockToken).Return(nil)

		err := service.DeleteReservation(id, mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}
func TestReservationService_GetReservationByID(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewReservationData(t)
	service := NewReservationService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := reservations.Reservation{}
		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		id := "1"

		claim := mockToken.Claims.(jwt.MapClaims)

		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("GetReservationByID", id, mockToken).Return(newData, nil)

		_, err := service.GetReservationByID(id, mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}
