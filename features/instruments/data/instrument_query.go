package data

import (
	"context"
	"project/features/instruments"
	"project/helper"

	"firebase.google.com/go/db"
)

type InstrumentData struct {
	db *db.Client
}

func NewInstrumentData(client *db.Client) instruments.InstrumentDataInterface {
	return &InstrumentData{
		db: client,
	}
}

func (id *InstrumentData) GetAllInstruments() (*instruments.InstrumentsMap, error) {
	ref := id.db.NewRef("instruments")
	var instruments instruments.InstrumentsMap
	if err := ref.Get(context.Background(), &instruments); err != nil {
		return nil, err
	}
	return &instruments, nil
}

func (id *InstrumentData) GetInstrumentByID(instrumentID string) (*instruments.Instruments, error) {
	ref := id.db.NewRef("instruments").Child(instrumentID)
	var instrument instruments.Instruments
	if err := ref.Get(context.Background(), &instrument); err != nil {
		return nil, err
	}
	return &instrument, nil
}

func (id *InstrumentData) AddInstrument(newData instruments.Instruments) (*instruments.Instruments, error) {
	ref := id.db.NewRef("instruments").Child(newData.IntrumentID)
	if err := ref.Set(context.Background(), &newData); err != nil {
		return nil, err
	}
	return &newData, nil
}

func (id *InstrumentData) UpdateInstrument(instrumentID string, newData instruments.Instruments) (*instruments.Instruments, error) {
	ref := id.db.NewRef("instruments").Child(instrumentID)
	res, err := helper.ToMap(newData)
	if err != nil {
		return nil, err
	}

	if err := ref.Update(context.Background(), res); err != nil {
		return nil, err
	}
	return &newData, nil
}

func (id *InstrumentData) DeleteInstrument(instrumentID string) error {
	ref := id.db.NewRef("instruments").Child(instrumentID)
	if err := ref.Delete(context.Background()); err != nil {
		return err
	}
	return nil
}
