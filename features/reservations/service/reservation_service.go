package service

import (
	"fmt"
	"project/features/reservations"
	"project/helper"

	"github.com/golang-jwt/jwt/v5"
)

type ReservationService struct {
	d reservations.ReservationData
	j helper.JWTInterface
}

func NewReservationService(data reservations.ReservationData, jwt helper.JWTInterface) *ReservationService {
	return &ReservationService{
		d: data,
		j: jwt,
	}
}

func (rs *ReservationService) GetAllReservations() (*reservations.AllReservations, error) {
	res, err := rs.d.GetAllReservations()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rs *ReservationService) GetReservationsByUsername(username string, token *jwt.Token) (map[string]any, error) {
	id, role := rs.j.ExtractToken(token)
	fmt.Println("id : ", id, ", role : ", role)
	res, err := rs.d.GetReservationsByUsername(username)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rs *ReservationService) AddReservation(newData reservations.Reservation) (*reservations.Reservation, error) {
	res, err := rs.d.AddReservation(newData)
	if err != nil {
		return nil, err
	}
	return res, nil
}
