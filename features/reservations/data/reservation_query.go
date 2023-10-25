package data

import (
	"context"
	"project/features/reservations"

	"firebase.google.com/go/db"
)

type ReservationData struct {
	db *db.Client
}

func NewReservationData(client db.Client) reservations.ReservationData {
	return &ReservationData{
		db: &client,
	}
}

func (rd *ReservationData) AddReservation(resData reservations.Reservation) (*reservations.Reservation, error) {
	ref := rd.db.NewRef("users/" + resData.Username + "/reservations").Child(resData.ReservationID)
	if err := ref.Set(context.Background(), resData); err != nil {
		return nil, err
	}

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

func (rd *ReservationData) GetReservationsByUsername(uname string) (*reservations.Reservation, error) {
	ref := rd.db.NewRef("reservations/").OrderByChild("username").EqualTo(uname)
	var res reservations.Reservation
	if err := ref.Get(context.Background(), &res); err != nil {
		return nil, err
	}
	return &res, nil
}
