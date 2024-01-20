package web

import "github.com/a-fandy/finan/model/entity"

type UserRequest struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required,max=15"`
}

func UserRequestToEntity(userRequest UserRequest) entity.User {
	return entity.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		NoHp:     userRequest.NoHp,
	}
}
