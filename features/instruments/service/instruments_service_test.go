package service

import (
	"errors"
	"project/features/instruments"
	"project/features/instruments/mocks"
	helperMock "project/helper/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestInstrumentService_GetAllInstruments(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewInstrumentDataInterface(t)
	service := NewInstrumentService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := map[string]map[string]any{
			"name": {
				"name": "fender stratocaster",
			},
		}

		// mockToken := &jwt.Token{
		// 	Claims: jwt.MapClaims{"role": "admin"},
		// }
		// jwtt.On("ExtractToken", mockToken).Return("id", "role")
		data.On("GetAllInstruments").Return(newData, nil)

		page := 1
		pageSize := 5
		_, err := service.GetAllInstruments(page, pageSize)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}
func TestInstrumentervice_GetInstrumentByID(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewInstrumentDataInterface(t)
	service := NewInstrumentService(data, jwtt)

	t.Run("Failed data retrieval", func(t *testing.T) {
		id := "1"
		data.On("GetInstrumentByID", id).Return(nil, errors.New("no data found"))
		_, err := service.GetInstrumentByID(id)
		assert.NotNil(t, err)
	})
	data.AssertExpectations(t)
}

func TestInstrumentService_AddInstrument(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewInstrumentDataInterface(t)
	service := NewInstrumentService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := instruments.Instruments{
			Name: "fender stratocaster",
			Type: "guitar",
			Year: 1959,
		}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{
				"role": "admin",
				"id":   "arya",
			},
		}
		claim := mockToken.Claims.(jwt.MapClaims)
		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("AddInstrument", newData, mockToken).Return(newData, nil)

		_, err := service.AddInstrument(newData, mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestInstrument_Delete(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewInstrumentDataInterface(t)
	service := NewInstrumentService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{
				"role": "admin",
				"id":   "arya",
			},
		}
		id := "1"
		claim := mockToken.Claims.(jwt.MapClaims)
		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("IsInstrumentExist", id).Return(true).Once()
		data.On("DeleteInstrument", id, mockToken).Return(nil)

		err := service.DeleteInstrument(id, mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestInstrument_Update(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewInstrumentDataInterface(t)
	service := NewInstrumentService(data, jwtt)

	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := instruments.Instruments{
			Name: "fender stratocaster",
			Type: "guitar",
			Year: 1959,
		}
		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{
				"role": "admin",
				"id":   "arya",
			},
		}
		id := "1"
		claim := mockToken.Claims.(jwt.MapClaims)
		jwtt.On("ExtractToken", mockToken).Return("id", claim["role"])
		data.On("IsInstrumentExist", id).Return(true).Once()
		data.On("UpdateInstrument", id, mockToken).Return(nil)

		_, err := service.UpdateInstrument(id, newData, mockToken)
		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}
