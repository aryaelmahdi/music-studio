package handler

import "project/features/rooms"

type UpdateRoomInput struct {
	Instruments rooms.Instruments `json:"instrument"`
	Price       int               `json:"price"`
}
