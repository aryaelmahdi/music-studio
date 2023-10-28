package service

import (
	"errors"
	"fmt"
	"project/features/reservations"
	"project/helper"
	"strings"

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

func (rs *ReservationService) GetAllReservations(token *jwt.Token) (*reservations.AllReservations, error) {
	_, role := rs.j.ExtractToken(token)
	if role != "admin" {
		return nil, errors.New("Unauthorized user")
	}
	res, err := rs.d.GetAllReservations()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rs *ReservationService) GetReservationsByUsername(token *jwt.Token) (map[string]any, error) {
	username, _ := rs.j.ExtractToken(token)
	res, err := rs.d.GetReservationsByUsername(fmt.Sprint(username))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rs *ReservationService) AddReservation(newData reservations.Reservation, token *jwt.Token) (*reservations.Reservation, error) {
	username, _ := rs.j.ExtractToken(token)
	newData.Username = fmt.Sprint(username)
	res, err := rs.d.AddReservation(newData)
	if err != nil {
		if strings.Contains(err.Error(), "room reserved") {
			return nil, err
		}
		return nil, err
	}
	return res, nil
}

func (rs *ReservationService) UpdateReservation(newData reservations.Reservation, token *jwt.Token) (*reservations.Reservation, error) {
	username, _ := rs.j.ExtractToken(token)
	newData.Username = fmt.Sprint(username)
	dataMap, err := helper.ToMap(newData)
	if err != nil {
		return nil, err
	}

	res, err := rs.d.UpdateReservation(dataMap)
	if err != nil {
		return nil, err
	}
	return res, nil
}
