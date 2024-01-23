package profile

import userEntity "github.com/a-fandy/finan/api/v1/user"

type UserRequestUpdate struct {
	Password string `json:"password" validate:"max=50"`
	Email    string `json:"email" validate:"required,email,max=200"`
	Name     string `json:"name" validate:"required,max=200"`
	NoHp     string `json:"no_hp" validate:"required,max=15"`
}

func UserRequestUpdateToEntity(userRequest UserRequestUpdate) userEntity.User {
	var user userEntity.User
	user.Email = userRequest.Email
	user.NoHp = userRequest.NoHp
	user.Name = userRequest.Name
	if userRequest.Password != "" {
		user.Password = userRequest.Password
	}
	return user
}
