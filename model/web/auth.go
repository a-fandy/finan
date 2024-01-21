package web

type LoginRequest struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
