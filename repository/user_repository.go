package repository

import (
	"context"

	"github.com/a-fandy/finan/model/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User) entity.User
	Update(ctx context.Context, user entity.User) entity.User
	Delete(ctx context.Context, user entity.User)
	FindById(ctx context.Context, id uint64) (entity.User, error)
	FindAll(ctx context.Context) []entity.User
}
