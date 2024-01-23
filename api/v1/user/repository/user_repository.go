package repository

import (
	"context"

	userEntity "github.com/a-fandy/finan/api/v1/user"
)

type UserRepository interface {
	Insert(ctx context.Context, user userEntity.User) userEntity.User
	Update(ctx context.Context, user userEntity.User, userUpdate userEntity.User) userEntity.User
	Delete(ctx context.Context, user userEntity.User)
	FindById(ctx context.Context, id uint64) (userEntity.User, error)
	FindByEmail(ctx context.Context, email string) (userEntity.User, error)
	FindAll(ctx context.Context) []userEntity.User
}
