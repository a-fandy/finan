package impl

import (
	"context"
	"errors"

	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/model/entity"
	"github.com/a-fandy/finan/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB}
}

func (repository UserRepositoryImpl) Insert(ctx context.Context, user entity.User) entity.User {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	exception.PanicIfError(err)
	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, user entity.User) entity.User {
	err := repository.DB.WithContext(ctx).Where("id = ?", user.Id).Updates(&user).Error
	exception.PanicIfError(err)
	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, user entity.User) {
	err := repository.DB.WithContext(ctx).Delete(&user).Error
	exception.PanicIfError(err)
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User
	if err := repository.DB.WithContext(ctx).Unscoped().Where("id = ?", id).First(&user).Error; err != nil {
		return entity.User{}, errors.New("User Not Found")
	}
	return user, nil
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context) []entity.User {
	var users []entity.User
	repository.DB.WithContext(ctx).Find(&users)
	return users
}

func (repository UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := repository.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, errors.New("User Not Found")
	}
	return user, nil
}
