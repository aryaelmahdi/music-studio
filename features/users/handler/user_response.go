package handler

type RegisterResponse struct {
	Nama string `json:"nama"`
	HP   string `json:"hp"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    any    `json:"token"`
}
