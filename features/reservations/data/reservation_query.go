package data

import (
	"context"
	"errors"
	"fmt"
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

	ref := rd.db.NewRef("reservations").Child(resData.RoomID)
	if err := ref.Set(context.Background(), resData); err != nil {
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

func (rd *ReservationData) UpdateReservation(newData map[string]interface{}) (*reservations.Reservation, error) {
	isValid, room, err := rd.isRoomValid(newData["room_id"].(string))
	if !isValid || err != nil {
		return nil, err
	}

	newData["price"] = room["price"]

	isUserValid, err := rd.isUserValid(newData["username"].(string), newData["reservation_id"].(string))
	if !isUserValid || err != nil {
		return nil, err
	}

	dateExists, roomExists := rd.isExist(newData["date"].(string), newData["room_id"].(string))
	if dateExists && roomExists {
		return nil, errors.New("room reserved")
	}

	ref := rd.db.NewRef("reservations").Child(newData["reservation_id"].(string))
	var res reservations.Reservation

	if err := ref.Update(context.Background(), newData); err != nil {
		return nil, err
	}

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

func (rd *ReservationData) DeleteReservation(id string) error {
	ref := rd.db.NewRef("reservations").Child(id)
	if err := ref.Delete(context.Background()); err != nil {
		return err
	}
	return nil
}

func (rd *ReservationData) GetReservationByID(id string) (*reservations.Reservation, error) {
	ref := rd.db.NewRef("reservations").Child(id)
	var reservation reservations.Reservation
	if err := ref.Get(context.Background(), &reservation); err != nil {
		return nil, errors.New("invalid reservation id")
	}
	return &reservation, nil
}

func (rd *ReservationData) isRoomValid(roomID string) (bool, map[string]any, error) {
	ref := rd.db.NewRef("rooms").Child(roomID)
	var room map[string]any
	if err := ref.Get(context.Background(), &room); err != nil {
		log.Fatal("Cannot get room")
		return false, room, err
	}
	return true, room, nil
}

func (rd *ReservationData) isExist(reservationDate string, roomID string) (bool, bool) {
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
			if dateStr, ok := date.(string); ok && dateStr == reservationDate {
				foundData[key] = data
				dateExists = true
				if room, ok := data["room_id"].(string); ok && room == roomID {
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

func (rd *ReservationData) isUserValid(username string, reservation_id string) (bool, error) {
	ref := rd.db.NewRef("reservations").Child(reservation_id)
	var res reservations.Reservation
	if err := ref.Get(context.Background(), &res); err != nil {
		return false, err
	}
	fmt.Print("res : ", res)
	if res.Username == username {
		return true, nil
	}
	return false, errors.New("username does not macthed!")
}
