package handler

import (
	"project/features/instruments"
)

type UpdateRoomInput struct {
	Instruments instruments.InstrumentsMap `json:"instrument"`
	Price       int                        `json:"price"`
}
