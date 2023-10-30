package handler

type ReservationRequest struct {
	ReservationID string `json:"reservation_id"`
	RoomID        string `json:"room_id"`
	PaymentStatus string `json:"payment_status"`
	Date          string `json:"date"`
}
