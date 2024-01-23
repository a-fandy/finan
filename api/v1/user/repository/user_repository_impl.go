package repository

import (
	"context"

	userEntity "github.com/a-fandy/finan/api/v1/user"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB}
}

func (repository UserRepositoryImpl) Insert(ctx context.Context, user userEntity.User) userEntity.User {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	exception.PanicIfError(err)
	repository.DB.WithContext(ctx).Model(&user).Update("refferal_code", helper.GenerateRefferalCode(ctx, repository.DB))
	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, user userEntity.User, userUpdate userEntity.User) userEntity.User {
	err := repository.DB.WithContext(ctx).Model(&user).Updates(&userUpdate).Error
	exception.PanicIfError(err)
	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, user userEntity.User) {
	err := repository.DB.WithContext(ctx).Delete(&user).Error
	exception.PanicIfError(err)
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, id uint64) (userEntity.User, error) {
	var dataUser userEntity.User
	if err := repository.DB.WithContext(ctx).Unscoped().Where("id = ?", id).First(&dataUser).Error; err != nil {
		return userEntity.User{}, exception.NotFoundError{Message: "User Not Found"}
	}
	return dataUser, nil
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context) []userEntity.User {
	var users []userEntity.User
	repository.DB.WithContext(ctx).Find(&users)
	return users
}

func (repository UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (userEntity.User, error) {
	var dataUser userEntity.User
	if err := repository.DB.WithContext(ctx).Where("email = ?", email).First(&dataUser).Error; err != nil {
		return userEntity.User{}, exception.NotFoundError{Message: "User Not Found"}
	}
	return dataUser, nil
}
