package repository

import (
	"context"

	incomeEntity "github.com/a-fandy/finan/api/v1/income"
	userEntity "github.com/a-fandy/finan/api/v1/user"
	"github.com/a-fandy/finan/exception"
	"gorm.io/gorm"
)

type IncomeRepositoryImpl struct {
	*gorm.DB
}

func NewIncomeRepository(DB *gorm.DB) IncomeRepository {
	return &IncomeRepositoryImpl{DB: DB}
}

func (incomeRepository IncomeRepositoryImpl) Insert(ctx context.Context, user userEntity.User, income incomeEntity.Income) incomeEntity.Income {
	err := incomeRepository.DB.WithContext(ctx).Model(&user).Association("Incomes").Append(append(user.Incomes, income))
	exception.PanicIfError(err)
	return income
}
