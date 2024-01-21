package web

type UserRequest struct {
	Password string `json:"password" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required,max=15"`
}

type UserRequestUpdate struct {
	Password string `json:"password" validate:"max=50"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required,max=15"`
}
