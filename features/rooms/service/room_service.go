package service

import (
	"errors"
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
	res, err := rs.d.AddRoom(newRoom)
	if err != nil {
		return nil, errors.New("Cannot Add Room")
	}
	return res, nil
}

func (rs *RoomService) DeleteRoom(roomID string) (string, error) {
	err := rs.d.DeleteRoom(roomID)
	if err != nil {
		return "", errors.New("Cannot delete room")
	}
	return roomID, nil
}

func (rs *RoomService) GetAllRooms() (*rooms.RoomMap, error) {
	res, err := rs.d.GetAllRooms()
	if err != nil {
		return nil, errors.New("Cannot get rooms data")
	}
	return res, nil
}

func (rs *RoomService) GetRoomByID(roomID string) (*rooms.Rooms, error) {
	res, err := rs.d.GetRoomByID(roomID)
	if err != nil {
		return nil, errors.New("Cannot get room data")
	}
	return res, nil
}

func (rs *RoomService) UpdateRoom(roomID string, updatedRoom rooms.Rooms) (*rooms.Rooms, error) {
	res, err := rs.d.UpdateRoom(roomID, updatedRoom)
	if err != nil {
		return nil, errors.New("Cannot update room")
	}
	return res, nil
}
