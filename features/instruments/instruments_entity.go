package instruments

import "github.com/labstack/echo/v4"

type Instruments struct {
	IntrumentID string `json:"instrument_id"`
	Name        string `json:"name"`
	Year        int    `json:"year"`
	Type        string `json:"type"`
}

type InstrumentsMap map[string]string

type AllInstruments struct {
	Data map[string]Instruments `json:"instruments"`
}

type InstrumentDataInterface interface {
	GetAllInstruments() (map[string]interface{}, error)
	GetInstrumentByID(id string) (*Instruments, error)
	AddInstrument(newData Instruments) (*Instruments, error)
	DeleteInstrument(id string) error
	UpdateInstrument(id string, newData Instruments) (*Instruments, error)
}

type InstrumentService interface {
	GetAllInstruments() (map[string]interface{}, error)
	GetInstrumentByID(id string) (*Instruments, error)
	AddInstrument(newData Instruments) (*Instruments, error)
	DeleteInstrument(id string) error
	UpdateInstrument(id string, newData Instruments) (*Instruments, error)
}

type InstrumentHandler interface {
	GetAllInstruments() echo.HandlerFunc
	GetInstrumentByID() echo.HandlerFunc
	AddInstrument() echo.HandlerFunc
	UpdateInstrument() echo.HandlerFunc
	DeleteInstrument() echo.HandlerFunc
}
