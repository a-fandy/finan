package login

import (
	userEntity "github.com/a-fandy/finan/api/v1/user"
	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/helper"
)

type LoginRequest struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func AuthToLoginResponse(user userEntity.User, conf config.Config) LoginResponse {
	return LoginResponse{Token: helper.GenerateJwtToken(user, conf.GetPrivateKey())}
}
