package helper

import (
	"context"
	"fmt"

	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/model/entity"
	"gorm.io/gorm"
)

func GenerateRefferalCode(ctx context.Context, DB *gorm.DB) string {
	var refFormat entity.RefFormat
	err := DB.WithContext(ctx).Where("format_id = ?", "referral_code").First(&refFormat).Error
	exception.PanicIfError(err)
	refFormat.Counter += 1
	err = DB.WithContext(ctx).Where("format_id = ?", "referral_code").Updates(&refFormat).Error
	exception.PanicIfError(err)
	code := fmt.Sprintf("%v%0*d", refFormat.ExtraField, refFormat.CounterPadLength, refFormat.Counter)
	return code
}
