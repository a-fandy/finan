package register

import (
	userEntity "github.com/a-fandy/finan/api/v1/user"
	"github.com/a-fandy/finan/helper"
)

type RegisterRequest struct {
	Password string `json:"password" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email,max=200"`
	Name     string `json:"name" validate:"required,max=200"`
	NoHp     string `json:"no_hp" validate:"required,max=15"`
}

func RegisterRequestToEntity(userRequest RegisterRequest) userEntity.User {
	return userEntity.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: helper.HashingPassword(userRequest.Password),
		NoHp:     userRequest.NoHp,
		Role:     "user",
		Status:   true,
		Username: helper.GenerateRandomString(5),
	}
}
