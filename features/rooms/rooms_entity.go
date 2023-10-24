package rooms

import "github.com/labstack/echo/v4"

type Instruments map[string]string

type Rooms struct {
	RoomID      string      `json:"room_id"`
	Instruments Instruments `json:"instrument"`
	Price       int         `json:"price"`
}

type RoomDataInterface interface {
	AddRoom(newRoom Rooms) (*Rooms, error)
	DeleteRoom(roomID string) error
	GetAllRooms() ([]Rooms, error)
	GetRoomByID(roomID string) (*Rooms, error)
	UpdateRoom(roomID string, updatedRoom Rooms) (*Rooms, error)
}

type RoomService interface {
	AddRoom(newRoom Rooms) (*Rooms, error)
	DeleteRoom(roomID string) (string, error)
	GetAllRooms() ([]Rooms, error)
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
