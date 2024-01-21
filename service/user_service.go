package service

import (
	"context"

	"github.com/a-fandy/finan/model/web"
)

type UserService interface {
	Create(ctx context.Context, userRequest web.UserRequest) web.UserResponse
	Update(ctx context.Context, userRequest web.UserRequestUpdate, id string) web.UserResponse
	Delete(ctx context.Context, id uint64)
	FindById(ctx context.Context, id string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
