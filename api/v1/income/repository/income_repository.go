package repository

import (
	"context"

	incomeEntity "github.com/a-fandy/finan/api/v1/income"
	userEntity "github.com/a-fandy/finan/api/v1/user"
)

type IncomeRepository interface {
	Insert(ctx context.Context, user userEntity.User, income incomeEntity.Income) incomeEntity.Income
	// Update(ctx context.Context, income entity.Income, incomeUpdate entity.Income) entity.Income
	// Delete(ctx context.Context, income entity.Income)
	// FindById(ctx context.Context, id uint64) (entity.Income, error)
	// FindAll(ctx context.Context) []entity.Income
}
