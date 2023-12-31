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

type Recommendation struct {
	Genre1 string `json:"genre1"`
	Genre2 string `json:"genre2"`
}

type RoomDataInterface interface {
	AddRoom(newRoom Rooms) (*Rooms, error)
	DeleteRoom(roomID string) error
	GetAllRooms() (map[string]map[string]interface{}, error)
	GetRoomByID(roomID string) (*Rooms, error)
	UpdateRoom(roomID string, updatedRoom Rooms) (*Rooms, error)
	AddRoomInstrument(roomId string, instrumentData instruments.RoomInstrument) (any, error)
	FilterRoomByPrice(price int) (map[string]map[string]interface{}, error)
	IsRoomExist(roomID string) bool
	GetBookedRooms() (map[string]map[string]interface{}, error)
	GetRecommendation(genre1 string, genre2 string, message string) (string, error)
}

type RoomService interface {
	AddRoom(newRoom Rooms, token *jwt.Token) (*Rooms, error)
	DeleteRoom(roomID string, token *jwt.Token) (any, error)
	GetAllRooms(page int, pageSize int) ([]map[string]any, error)
	GetRoomByID(roomID string) (*Rooms, error)
	UpdateRoom(roomID string, updatedRoom Rooms, token *jwt.Token) (*Rooms, error)
	AddRoomInstrument(roomId string, instrumentData instruments.RoomInstrument, token *jwt.Token) (any, error)
	FilterRoomByPrice(price int, page int, pageSize int) ([]map[string]any, error)
	GetBookedRooms(page int, pageSize int) ([]map[string]any, error)
	GetRecommendation(genre1 string, genre2 string) (any, error)
}

type RoomHandler interface {
	AddRoom() echo.HandlerFunc
	DeleteRoom() echo.HandlerFunc
	GetAllRooms() echo.HandlerFunc
	GetRoomByID() echo.HandlerFunc
	UpdateRoom() echo.HandlerFunc
	AddRoomInstrument() echo.HandlerFunc
	GetBookedRooms() echo.HandlerFunc
	GetRecommendation() echo.HandlerFunc
}
