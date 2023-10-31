package service

import (
	"errors"
	"fmt"
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

func (rs *RoomService) GetAllRooms() (map[string]any, error) {
	fmt.Println("masuk all")
	res, err := rs.d.GetAllRooms()
	if err != nil {
		return nil, errors.New("Cannot get rooms data")
	}
	if len(res) == 0 {
		return nil, errors.New("no data found")
	}
	return res, nil
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

func (rs *RoomService) FilterRoomByPrice(price int, page int, pageSize int) (map[string]any, error) {
	fmt.Println("masuk filter")
	res, err := rs.d.FilterRoomByPrice(price)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("No data")
	}
	paginatedRes := helper.PaginateMap(res, page, pageSize)
	return paginatedRes, nil
}
