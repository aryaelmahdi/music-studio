package data

import (
	"context"
	"errors"
	"fmt"
	"project/features/rooms"
	"project/helper"

	"firebase.google.com/go/db"
)

type RoomData struct {
	db *db.Client
}

func NewRoomData(client *db.Client) rooms.RoomDataInterface {
	return &RoomData{
		db: client,
	}
}

func (rd *RoomData) GetAllRooms() (*rooms.RoomMap, error) {
	ref := rd.db.NewRef("rooms")
	var rooms rooms.RoomMap
	if err := ref.Get(context.Background(), &rooms); err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (rd *RoomData) GetRoomByID(roomID string) (*rooms.Rooms, error) {
	ref := rd.db.NewRef("rooms").Child(roomID)
	var room rooms.Rooms
	if err := ref.Get(context.Background(), &room); err != nil {
		return nil, err
	}
	return &room, nil
}

func (rd *RoomData) AddRoom(newRoom rooms.Rooms) (*rooms.Rooms, error) {
	if roomExists := rd.isRoomExist(newRoom.RoomID); roomExists {
		return nil, errors.New("room already exists")
	}
	ref := rd.db.NewRef("rooms").Child(newRoom.RoomID)
	if err := ref.Set(context.Background(), newRoom); err != nil {
		return nil, err
	}
	return &newRoom, nil
}

func (rd *RoomData) DeleteRoom(roomID string) error {
	ref := rd.db.NewRef("rooms").Child(roomID)
	if err := ref.Delete(context.Background()); err != nil {
		return err
	}
	return nil
}

func (rd *RoomData) UpdateRoom(roomID string, updatedRoom rooms.Rooms) (*rooms.Rooms, error) {
	ref := rd.db.NewRef("rooms").Child(roomID)
	res, err := helper.ToMap(updatedRoom)
	if err != nil {
		return nil, err
	}

	if err := ref.Update(context.Background(), res); err != nil {
		return nil, err
	}
	return &updatedRoom, nil
}

func (rd *RoomData) isRoomExist(roomID string) bool {
	ref := rd.db.NewRef("rooms").Child(roomID)
	var room map[string]any
	ref.Get(context.Background(), &room)
	fmt.Println("room : ", room)
	if room == nil {
		return false
	}
	return true
}
