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
	user := web.UserRequestToEntity(userRequest)
	user.Role = "user"
	user.Status = true
	user = service.UserRepository.Insert(ctx, user)
	return web.UserEntityToResponse(user)
}

func (service UserServiceImpl) Update(ctx context.Context, userRequest web.UserRequest, id uint64) web.UserResponse {
	helper.Validate(userRequest)
	user := web.UserRequestToEntity(userRequest)
	user.Id = id
	user = service.UserRepository.Update(ctx, user)
	return web.UserEntityToResponse(user)

}

func (service UserServiceImpl) Delete(ctx context.Context, id uint64) {
	users, err := service.UserRepository.FindById(ctx, id)
	exception.PanicIfError(exception.NotFoundError{Message: err.Error()})
	service.UserRepository.Delete(ctx, users)
}

func (service UserServiceImpl) FindById(ctx context.Context, id uint64) web.UserResponse {
	user, err := service.UserRepository.FindById(ctx, id)
	exception.PanicIfError(exception.NotFoundError{Message: err.Error()})
	return web.UserEntityToResponse(user)
}

func (service UserServiceImpl) FindAll(ctx context.Context) (userResponse []web.UserResponse) {
	users := service.UserRepository.FindAll(ctx)

	if len(users) == 0 {
		return []web.UserResponse{}
	}

	for _, user := range users {
		userResponse = append(userResponse, web.UserEntityToResponse(user))
	}

	return userResponse
}
