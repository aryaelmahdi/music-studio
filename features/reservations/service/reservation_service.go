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
	if len(res) == 0 {
		return nil, errors.New("no data found")
	}
	return res, nil
}

func (rs *ReservationService) AddReservation(newData reservations.Reservation, token *jwt.Token) (*reservations.Reservation, error) {
	username, _ := rs.j.ExtractToken(token)
	newData.Username = fmt.Sprint(username)
	newData.ReservationID = newData.Date + newData.RoomID
	newData.PaymentStatus = "not finished"
	res, err := rs.d.AddReservation(newData)
	if err != nil {
		if strings.Contains(err.Error(), "room reserved") {
			return nil, err
		}
		return nil, err
	}
	return res, nil
}

func (rs *ReservationService) UpdateReservation(id string, newData reservations.Reservation, token *jwt.Token) (*reservations.Reservation, error) {
	dateErr := helper.VerifyCancelDate(newData.Date)
	if dateErr != nil {
		return nil, dateErr
	}

	username, _ := rs.j.ExtractToken(token)
	newData.Username = fmt.Sprint(username)
	newData.ReservationID = id
	dataMap, err := helper.ToMap(newData)
	if err != nil {
		return nil, err
	}

	roomExists, room, err := rs.d.IsRoomValid(id)
	if roomExists {
		newData.Price = room["price"]
	}

	res, err := rs.d.UpdateReservation(newData.ReservationID, dataMap)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rs *ReservationService) DeleteReservation(id string, token *jwt.Token) error {
	username, role := rs.j.ExtractToken(token)
	res, _ := rs.d.GetReservationByID(id)
	fmt.Println("res username :", res.Username, "username :", username)
	if res.Username != username && role != "admin" {
		return errors.New("Unauthorized user")
	}
	if err := rs.d.DeleteReservation(id); err != nil {
		return err
	}
	return nil
}

func (rd *ReservationService) GetReservationByID(id string, token *jwt.Token) (*reservations.Reservation, error) {
	username, role := rd.j.ExtractToken(token)
	res, err := rd.d.GetReservationByID(id)
	if err != nil {
		return nil, err
	}
	if res.Username != username && role != "admin" {
		return nil, errors.New("Unauthorized user")
	}
	return res, nil
}
