package web

type UserResponse struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	NoHp     string `json:"no_hp"`
	Status   bool   `json:"status"`
}
