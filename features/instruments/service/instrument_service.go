package service

import (
	"errors"
	"project/features/instruments"
	"project/helper"
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

func (is *InstrumentService) GetAllInstruments() (map[string]interface{}, error) {
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
	return res, nil
}

func (is *InstrumentService) AddInstrument(newData instruments.Instruments) (*instruments.Instruments, error) {
	res, err := is.d.AddInstrument(newData)
	if err != nil {
		return nil, errors.New("Canont Add instrument " + err.Error())
	}
	return res, nil
}

func (is *InstrumentService) DeleteInstrument(id string) error {
	err := is.d.DeleteInstrument(id)
	if err != nil {
		return errors.New("Cannot delete instrument " + err.Error())
	}
	return nil
}

func (is *InstrumentService) UpdateInstrument(id string, newData instruments.Instruments) (*instruments.Instruments, error) {
	res, err := is.d.UpdateInstrument(id, newData)
	if err != nil {
		return nil, errors.New("Cannot update instrument" + err.Error())
	}
	return res, nil
}
