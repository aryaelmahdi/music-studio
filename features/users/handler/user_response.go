package handler

type RegisterResponse struct {
	Username string `json:"username"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    any    `json:"token"`
}
