package rooms

type Instruments map[string]string

type Rooms struct {
	RoomID      string      `json:"room_id"`
	Instruments Instruments `json:"instruments"`
	Price       int         `json:"price"`
}

type RoomDataInterface interface {
}
