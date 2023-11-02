package service

import (
	"errors"
	"project/features/instruments"
	"project/features/rooms"
	"project/features/rooms/mocks"
	helperMock "project/helper/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestRoomService_GetAllRooms(t *testing.T) {
	jwt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwt)
	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := map[string]map[string]any{
			"room_id": {
				"name": "ruangA",
			},
		}

		data.On("GetAllRooms").Return(newData, nil)

		page := 1
		pageSize := 10

		_, err := service.GetAllRooms(page, pageSize)

		assert.Nil(t, err)
	})
	t.Run("Fail data retrieval", func(t *testing.T) {

		newData := map[string]map[string]any{
			"room_id": {
				"name": "ruangA",
			},
		}

		data.On("GetAllRooms").Return(newData, nil)

		page := 0
		pageSize := 0

		result, err := service.GetAllRooms(page, pageSize)

		assert.Nil(t, result)
		assert.Equal(t, errors.New("Invalid page or pageSize"), err)
	})
	data.AssertExpectations(t)
}

func TestRoomService_GetRoomByID(t *testing.T) {
	jwt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwt)
	t.Run("Successful data retrieval", func(t *testing.T) {
		expectedRoom := &rooms.Rooms{
			RoomID: "1",
			Price:  20,
		}
		roomID := "1"
		data.On("GetRoomByID", roomID).Return(expectedRoom, nil)

		result, err := service.GetRoomByID(roomID)

		assert.Nil(t, err)
		assert.Equal(t, expectedRoom, result)
	})
	data.AssertExpectations(t)
}

func TestRoomService_UpdateRoom(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwtt)
	t.Run("Successful room update", func(t *testing.T) {
		roomID := "1"
		updatedRoom := rooms.Rooms{
			RoomID: "1",
			Price:  50,
		}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		data.On("UpdateRoom", roomID, updatedRoom, mockToken).Return(updatedRoom, nil)
		result, err := service.UpdateRoom(roomID, updatedRoom, mockToken)

		assert.Nil(t, err)
		assert.Equal(t, updatedRoom, *result)
	})
	data.AssertExpectations(t)
}

func TestRoomService_AddRoomInstrument(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwtt)
	t.Run("Successful instrument addition", func(t *testing.T) {
		roomID := "1"
		instrumentData := instruments.RoomInstrument{}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		jwtt.On("ExtractToken", mockToken).Return("id", "role")
		data.On("AddRoomInstrument", roomID, instrumentData, mockToken).Return(instrumentData, nil)
		_, err := service.AddRoomInstrument(roomID, instrumentData, mockToken)

		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestRoomData_AddRoom(t *testing.T) {
	jwtt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwtt)
	t.Run("Successful room addition", func(t *testing.T) {
		newRoom := rooms.Rooms{
			RoomID: "101",
			Price:  100,
		}

		secondRoom := rooms.Rooms{}

		mockToken := &jwt.Token{
			Claims: jwt.MapClaims{"role": "admin"},
		}
		jwtt.On("ExtractToken", mockToken).Return("id", "role")
		data.On("AddRoom", newRoom, mockToken).Return(secondRoom, nil)
		_, err := service.AddRoom(newRoom, mockToken)

		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestRoomService_FilterRooms(t *testing.T) {
	jwt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwt)
	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := map[string]map[string]any{
			"room_id": {
				"name": "ruangA",
			},
		}

		price := 20
		data.On("FilterRoomByPrice", price).Return(newData, nil)
		page := 1
		pageSize := 10

		_, err := service.FilterRoomByPrice(price, page, pageSize)

		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestRoomService_GetBookedRooms(t *testing.T) {
	jwt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwt)
	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := map[string]map[string]any{
			"room_id": {
				"name": "ruangA",
			},
		}

		data.On("GetBookedRooms").Return(newData, nil)
		page := 1
		pageSize := 10

		_, err := service.GetBookedRooms(page, pageSize)

		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}

func TestRoomService_GetRecommendation(t *testing.T) {
	jwt := helperMock.NewJWTInterface(t)
	data := mocks.NewRoomDataInterface(t)
	var service = NewRoomService(data, jwt)
	t.Run("Successful data retrieval", func(t *testing.T) {
		newData := map[string]map[string]any{
			"room_id": {
				"name": "ruangA",
			},
		}

		response := "based on data give..."

		data.On("GetAllRooms").Return(newData, nil)
		genre1 := "rock"
		genre2 := "metal"
		message := "i want to play rock and metal in a studio, these are the rooms data the gave me : {\"room_id\":{\"name\":\"ruangA\"}}. \n based on the instruments on each rooms, which one is suitable for my purpose?  for the room names, don't include the / into your response, don't need to make a new line, don't need to include backslash to point out the data, slash or any symbols, just straight up room names and its instruments like 'ruangH - fender stratocaster. reasons....' in a single paragraph, not like \\ruangB - fender stratocaster\\"
		data.On("GetRecommendation", genre1, genre2, message).Return(response, nil)

		_, err := service.GetRecommendation(genre1, genre2)

		assert.Nil(t, err)
	})
	data.AssertExpectations(t)
}
