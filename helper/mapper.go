package helper

import (
	"github.com/a-fandy/finan/model/entity"
	"github.com/a-fandy/finan/model/web"
)

func UserRequestToEntity(userRequest web.UserRequest) entity.User {
	Validate(userRequest)
	return entity.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: HashingPassword(userRequest.Password),
		NoHp:     userRequest.NoHp,
	}
}

func UserEntityToResponse(user entity.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		NoHp:     user.NoHp,
		Status:   user.Status,
	}
}

func AuthToLoginResponse(user entity.User) web.LoginResponse {
	return web.LoginResponse{Token: "masuk"}
}
