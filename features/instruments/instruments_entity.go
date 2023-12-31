package instruments

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Instruments struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Year int    `json:"year"`
}

type RoomInstrument struct {
	Guitar   string `json:"guitar"`
	Bass     string `json:"bass"`
	Drum     string `json:"drum"`
	Keyboard string `json:"keyboard"`
}

type InstrumentsMap map[string]interface{}

type InstrumentDataInterface interface {
	GetAllInstruments() (map[string]map[string]interface{}, error)
	GetInstrumentByID(id string) (*Instruments, error)
	AddInstrument(newData Instruments) (*Instruments, error)
	DeleteInstrument(id string) error
	UpdateInstrument(id string, newData Instruments) (*Instruments, error)
	IsInstrumentExist(instrumentName string) bool
}

type InstrumentService interface {
	GetAllInstruments(page int, pageSize int) ([]map[string]any, error)
	GetInstrumentByID(id string) (*Instruments, error)
	AddInstrument(newData Instruments, token *jwt.Token) (*Instruments, error)
	DeleteInstrument(id string, token *jwt.Token) error
	UpdateInstrument(id string, newData Instruments, token *jwt.Token) (*Instruments, error)
}

type InstrumentHandler interface {
	GetAllInstruments() echo.HandlerFunc
	GetInstrumentByID() echo.HandlerFunc
	AddInstrument() echo.HandlerFunc
	UpdateInstrument() echo.HandlerFunc
	DeleteInstrument() echo.HandlerFunc
}
