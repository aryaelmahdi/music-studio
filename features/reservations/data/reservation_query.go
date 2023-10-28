package data

import (
	"context"
	"errors"
	"log"
	"project/features/reservations"
	"project/helper"

	"firebase.google.com/go/db"
)

type ReservationData struct {
	db *db.Client
}

func NewReservationData(client *db.Client) reservations.ReservationData {
	return &ReservationData{
		db: client,
	}
}

func (rd *ReservationData) AddReservation(resData reservations.Reservation) (*reservations.Reservation, error) {
	dateExists, roomExists := rd.isExist(resData.Date, resData.RoomID)
	if dateExists && roomExists {
		return nil, errors.New("room reserved")
	}

	roomValid, room, err := rd.isRoomValid(resData.RoomID)
	if !roomValid || err != nil {
		return nil, errors.New("room does not exist")
	}

	price, err := helper.ExtractPrice(room)
	if err != nil {
		return nil, err
	}
	resData.Price = price

	newRef, err := rd.db.NewRef("reservations").Push(context.Background(), &resData)
	if err != nil {
		return nil, err
	}

	if err := newRef.Set(context.Background(), resData); err != nil {
		return nil, err
	}

	return &resData, nil
}

func (rd *ReservationData) GetAllReservations() (*reservations.AllReservations, error) {
	ref := rd.db.NewRef("reservations")
	var res reservations.AllReservations
	if err := ref.Get(context.Background(), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (rd *ReservationData) GetReservationsByUsername(uname string) (map[string]any, error) {
	ref := rd.db.NewRef("reservations").OrderByChild("username").EqualTo(&uname)
	var res map[string]any
	if err := ref.Get(context.Background(), &res); err != nil {
		return nil, err
	}
	log.Print("res :", res)
	return res, nil
}

func (rd *ReservationData) isRoomValid(roomID string) (bool, map[string]any, error) {
	ref := rd.db.NewRef("rooms").Child(roomID)
	var room map[string]any
	if err := ref.Get(context.Background(), &room); err != nil {
		log.Fatal("Cannot get room")
		return false, room, err
	}
	log.Print("get", room)
	return true, room, nil
}

func (rd *ReservationData) isExist(date string, roomID string) (bool, bool) {
	reserved := map[string]map[string]interface{}{}
	ref := rd.db.NewRef("reservations")
	err := ref.Get(context.Background(), &reserved)
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	foundData := map[string]map[string]interface{}{}
	var roomExists bool = false
	var dateExists bool = false
	for key, data := range reserved {
		if date, exists := data["date"]; exists {
			if dateStr, ok := date.(string); ok && dateStr == date {
				foundData[key] = data
				dateExists = true
				if _, roomExists := data["room_id"].(string); roomExists {
					roomExists = true
					break
				}
			}
		}
	}
	if roomExists && dateExists {
		return true, true
	}
	return dateExists, roomExists
}
