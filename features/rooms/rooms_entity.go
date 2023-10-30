package rooms

import (
	"project/features/instruments"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Rooms struct {
	RoomID      string                     `json:"room_id"`
	Instruments instruments.RoomInstrument `json:"instrument"`
	Price       int                        `json:"price"`
}

type RoomMap map[string]interface{}

type RoomDataInterface interface {
	AddRoom(newRoom Rooms) (*Rooms, error)
	DeleteRoom(roomID string) error
	GetAllRooms() (*RoomMap, error)
	GetRoomByID(roomID string) (*Rooms, error)
	UpdateRoom(roomID string, updatedRoom Rooms) (*Rooms, error)
	AddRoomInstrument(roomId string, instrumentData instruments.RoomInstrument) (any, error)
	FilterRoomByPrice(price int) (map[string]any, error)
}

type RoomService interface {
	AddRoom(newRoom Rooms, token *jwt.Token) (*Rooms, error)
	DeleteRoom(roomID string, token *jwt.Token) (any, error)
	GetAllRooms() (*RoomMap, error)
	GetRoomByID(roomID string) (*Rooms, error)
	UpdateRoom(roomID string, updatedRoom Rooms) (*Rooms, error)
	AddRoomInstrument(roomId string, instrumentData instruments.RoomInstrument, token *jwt.Token) (any, error)
	FilterRoomByPrice(price int) (map[string]any, error)
}

type RoomHandler interface {
	AddRoom() echo.HandlerFunc
	DeleteRoom() echo.HandlerFunc
	GetAllRooms() echo.HandlerFunc
	GetRoomByID() echo.HandlerFunc
	UpdateRoom() echo.HandlerFunc
	AddRoomInstrument() echo.HandlerFunc
}
