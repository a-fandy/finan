package helper

import (
	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/model/entity"
	"github.com/a-fandy/finan/model/web"
)

func UserRequestToEntity(userRequest web.UserRequest) entity.User {
	return entity.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: HashingPassword(userRequest.Password),
		NoHp:     userRequest.NoHp,
	}
}

func UserRequestUpdateToEntity(userRequest web.UserRequestUpdate) entity.User {
	var user entity.User
	user.Email = userRequest.Email
	user.NoHp = userRequest.NoHp
	user.Name = userRequest.Name
	if userRequest.Password != "" {
		user.Password = userRequest.Password
	}
	return user
}

func UserEntityToResponse(user entity.User) web.UserResponse {
	return web.UserResponse{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		Name:         user.Name,
		NoHp:         user.NoHp,
		RefferalCode: user.RefferalCode,
		Status:       user.Status,
	}
}

func AuthToLoginResponse(user entity.User, conf config.Config) web.LoginResponse {
	return web.LoginResponse{Token: GenerateJwtToken(user, conf.GetPrivateKey())}
}
