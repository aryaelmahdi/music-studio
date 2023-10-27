package rooms

import (
	"project/features/instruments"

	"github.com/labstack/echo/v4"
)

type Rooms struct {
	RoomID      string                     `json:"room_id"`
	Instruments instruments.InstrumentsMap `json:"instrument"`
	Price       int                        `json:"price"`
}

type RoomMap map[string]interface{}

type RoomDataInterface interface {
	AddRoom(newRoom Rooms) (*Rooms, error)
	DeleteRoom(roomID string) error
	GetAllRooms() (*RoomMap, error)
	GetRoomByID(roomID string) (*Rooms, error)
	UpdateRoom(roomID string, updatedRoom Rooms) (*Rooms, error)
}

type RoomService interface {
	AddRoom(newRoom Rooms) (*Rooms, error)
	DeleteRoom(roomID string) (string, error)
	GetAllRooms() (*RoomMap, error)
	GetRoomByID(roomID string) (*Rooms, error)
	UpdateRoom(roomID string, updatedRoom Rooms) (*Rooms, error)
}

type RoomHandler interface {
	AddRoom() echo.HandlerFunc
	DeleteRoom() echo.HandlerFunc
	GetAllRooms() echo.HandlerFunc
	GetRoomByID() echo.HandlerFunc
	UpdateRoom() echo.HandlerFunc
}
