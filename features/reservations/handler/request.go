package handler

type ReservationRequest struct {
	ReservationID string `json:"reservation_id"`
	RoomID        string `json:"room_id"`
	Date          string `json:"date"`
}
