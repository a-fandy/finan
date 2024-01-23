package impl

import (
	"context"

	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/a-fandy/finan/model/web"
	"github.com/a-fandy/finan/repository"
	"github.com/a-fandy/finan/service"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) service.UserService {
	return &UserServiceImpl{UserRepository: *userRepository}
}

func (service UserServiceImpl) Create(ctx context.Context, userRequest web.UserRequest) web.UserResponse {
	helper.Validate(userRequest)
	user := helper.UserRequestToEntity(userRequest)
	user.Role = "user"
	user.Status = true
	user.Username = helper.GenerateRandomString(5)
	user = service.UserRepository.Insert(ctx, user)
	return helper.UserEntityToResponse(user)
}

func (service UserServiceImpl) Update(ctx context.Context, userRequest web.UserRequestUpdate, id string) web.UserResponse {
	helper.Validate(userRequest)
	user := helper.UserRequestUpdateToEntity(userRequest)
	user = service.UserRepository.Update(ctx, user, id)
	return helper.UserEntityToResponse(user)

}

func (service UserServiceImpl) Delete(ctx context.Context, id uint64) {
	users, err := service.UserRepository.FindById(ctx, id)
	exception.PanicIfError(err)
	service.UserRepository.Delete(ctx, users)
}

func (service UserServiceImpl) FindById(ctx context.Context, id string) web.UserResponse {
	user, err := service.UserRepository.FindByEmail(ctx, id)
	exception.PanicIfError(err)
	return helper.UserEntityToResponse(user)
}

func (service UserServiceImpl) FindAll(ctx context.Context) (userResponse []web.UserResponse) {
	users := service.UserRepository.FindAll(ctx)

	if len(users) == 0 {
		return []web.UserResponse{}
	}

	for _, user := range users {
		userResponse = append(userResponse, helper.UserEntityToResponse(user))
	}

	return userResponse
}
