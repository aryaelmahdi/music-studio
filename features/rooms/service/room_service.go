package service

import (
	"errors"
	"project/features/instruments"
	"project/features/rooms"
	"project/helper"

	"github.com/golang-jwt/jwt/v5"
)

type RoomService struct {
	d rooms.RoomDataInterface
	j helper.JWTInterface
}

func NewRoomService(data rooms.RoomDataInterface, jwt helper.JWTInterface) rooms.RoomService {
	return &RoomService{
		d: data,
		j: jwt,
	}
}

func (rs *RoomService) AddRoom(newRoom rooms.Rooms, token *jwt.Token) (*rooms.Rooms, error) {
	if _, role := rs.j.ExtractToken(token); role != "admin" {
		return nil, errors.New("Unauthorized user")
	}

	var currentRoom rooms.Rooms

	currentRoom.Price = newRoom.Price
	currentRoom.RoomID = newRoom.RoomID

	res, err := rs.d.AddRoom(currentRoom)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rs *RoomService) DeleteRoom(roomID string, token *jwt.Token) (any, error) {
	if _, role := rs.j.ExtractToken(token); role != "admin" {
		return nil, errors.New("Unauthorized user")
	}

	if roomExists := rs.d.IsRoomExist(roomID); !roomExists {
		return nil, errors.New("Invalid id")
	}

	err := rs.d.DeleteRoom(roomID)
	if err != nil {
		return nil, errors.New("Cannot delete room")
	}
	return roomID, nil
}

func (rs *RoomService) GetAllRooms(page int, pageSize int) ([]map[string]any, error) {
	res, err := rs.d.GetAllRooms()
	if err != nil {
		return nil, errors.New("Cannot get rooms data")
	}
	if len(res) == 0 {
		return nil, errors.New("no data found")
	}
	dataSlices := make([]map[string]any, 0)

	for _, data := range res {
		if _, exists := data["room_id"]; exists {
			dataSlices = append(dataSlices, data)
		}
	}

	paginatedRes, err := helper.Paginate(dataSlices, page, pageSize)
	if err != nil {
		return nil, err
	}
	return paginatedRes, nil
}

func (rs *RoomService) GetRoomByID(roomID string) (*rooms.Rooms, error) {
	res, err := rs.d.GetRoomByID(roomID)
	if err != nil {
		return nil, errors.New("Cannot get room data")
	}
	if res.RoomID == "" {
		return nil, errors.New("No data")
	}
	return res, nil
}

func (rs *RoomService) UpdateRoom(roomID string, updatedRoom rooms.Rooms, token *jwt.Token) (*rooms.Rooms, error) {
	_, role := rs.j.ExtractToken(token)
	if role != "admin" {
		return nil, errors.New("Unauthorized user")
	}
	if roomsExists := rs.d.IsRoomExist(roomID); !roomsExists {
		return nil, errors.New("invalid id")
	}
	res, err := rs.d.UpdateRoom(roomID, updatedRoom)
	if err != nil {
		return nil, errors.New("Cannot update room")
	}
	return res, nil
}

func (rs *RoomService) AddRoomInstrument(roomID string, instrumentData instruments.RoomInstrument, token *jwt.Token) (any, error) {
	_, role := rs.j.ExtractToken(token)
	if role != "admin" {
		return nil, errors.New("Unautorized user")
	}
	res, err := rs.d.AddRoomInstrument(roomID, instrumentData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rs *RoomService) FilterRoomByPrice(price int, page int, pageSize int) ([]map[string]any, error) {
	res, err := rs.d.FilterRoomByPrice(price)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("No data")
	}
	dataSlices := make([]map[string]any, 0)

	for _, data := range res {
		if _, exists := data["room_id"]; exists {
			dataSlices = append(dataSlices, data)
		}
	}

	paginatedRes, err := helper.Paginate(dataSlices, page, pageSize)
	if err != nil {
		return nil, err
	}
	return paginatedRes, nil
}

func (rs *RoomService) GetBookedRooms(page int, pageSize int) ([]map[string]any, error) {
	reserved, err := rs.d.GetBookedRooms()
	if err != nil {
		return nil, err
	}
	dataSlices := make([]map[string]interface{}, 0)

	for _, data := range reserved {
		date, dateExists := data["date"]
		room, roomExists := data["room_id"]

		if dateExists && roomExists {
			if dateStr, ok := date.(string); ok {
				if roomID, ok := room.(string); ok {
					newData := map[string]interface{}{
						"date":    dateStr,
						"room_id": roomID,
					}
					dataSlices = append(dataSlices, newData)
				}
			}
		}
	}
	res, err := helper.Paginate(dataSlices, page, pageSize)
	if err != nil {
		return nil, err
	}
	return res, nil
}
