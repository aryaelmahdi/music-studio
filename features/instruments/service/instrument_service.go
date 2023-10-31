package service

import (
	"errors"
	"project/features/instruments"
	"project/helper"

	"github.com/golang-jwt/jwt/v5"
)

type InstrumentService struct {
	d instruments.InstrumentDataInterface
	j helper.JWTInterface
}

func NewInstrumentService(data instruments.InstrumentDataInterface, jwt helper.JWTInterface) instruments.InstrumentService {
	return &InstrumentService{
		d: data,
		j: jwt,
	}
}

func (is *InstrumentService) GetAllInstruments() (*instruments.InstrumentsMap, error) {
	res, err := is.d.GetAllInstruments()
	if err != nil {
		return nil, errors.New("Cannot get Instruments data" + err.Error())
	}
	return res, nil
}

func (is *InstrumentService) GetInstrumentByID(id string) (*instruments.Instruments, error) {
	res, err := is.d.GetInstrumentByID(id)
	if err != nil {
		return nil, errors.New("Cannot get Intrument data " + err.Error())
	}
	if res.Name == "" {

	}
	return res, nil
}

func (is *InstrumentService) AddInstrument(newData instruments.Instruments, token *jwt.Token) (*instruments.Instruments, error) {
	_, role := is.j.ExtractToken(token)
	if role != "admin" {
		return nil, errors.New("Unauthorized user")
	}
	if newData.Name == "" && newData.Type == "" && newData.Year < 1750 {
		return nil, errors.New("Invalid input data")
	}
	if newData.Type != "guitar" && newData.Type != "drum" && newData.Type != "bass" && newData.Type != "keyboard" {
		return nil, errors.New("Invalid instrument type")
	}
	res, err := is.d.AddInstrument(newData)
	if err != nil {
		return nil, errors.New("Cannot Add instrument " + err.Error())
	}
	return res, nil
}

func (is *InstrumentService) DeleteInstrument(id string, token *jwt.Token) error {
	_, role := is.j.ExtractToken(token)
	if role != "admin" {
		return errors.New("Unauthorized user")
	}
	err := is.d.DeleteInstrument(id)
	if err != nil {
		return errors.New("Cannot delete instrument " + err.Error())
	}
	return nil
}

func (is *InstrumentService) UpdateInstrument(id string, newData instruments.Instruments, token *jwt.Token) (*instruments.Instruments, error) {
	_, role := is.j.ExtractToken(token)
	if role != "admin" {
		return nil, errors.New("Unauthorized user")
	}
	res, err := is.d.UpdateInstrument(id, newData)
	if err != nil {
		return nil, errors.New("Cannot update instrument" + err.Error())
	}
	return res, nil
}
