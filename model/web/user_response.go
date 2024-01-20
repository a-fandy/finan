package web

import "github.com/a-fandy/finan/model/entity"

type UserResponse struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	NoHp     string `json:"no_hp"`
	Status   bool   `json:"status"`
}

func UserEntityToResponse(user entity.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		NoHp:     user.NoHp,
		Status:   user.Status,
	}
}
