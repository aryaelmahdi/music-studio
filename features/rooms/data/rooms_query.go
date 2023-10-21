package data

import (
	"context"
	"project/features/rooms"

	"firebase.google.com/go/v4/db"
)

type RoomData struct {
	db *db.Client
}

func NewRoomData(client *db.Client) rooms.RoomDataInterface {
	return &RoomData{
		db: client,
	}
}

func (rd *RoomData) GetAllRooms() ([]rooms.Rooms, error) {
	ref := rd.db.NewRef("rooms")
	var rooms []rooms.Rooms
	if err := ref.Get(context.Background(), &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rd *RoomData) GetRoomByID(roomID string) (*rooms.Rooms, error) {
	ref := rd.db.NewRef("rooms").Child(roomID)
	var room rooms.Rooms
	if err := ref.Get(context.Background(), &room); err != nil {
		return nil, err
	}
	return &room, nil
}
