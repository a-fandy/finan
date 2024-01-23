package helper

import (
	"context"
	"fmt"

	"github.com/a-fandy/finan/exception"
	"gorm.io/gorm"
)

type RefFormat struct {
	FormatId         string `gorm:"type:VARCHAR(255)" json:"format_id"`
	Format           string `gorm:"type:VARCHAR(255)" json:"format"`
	Counter          int    `gorm:"type:INT" json:"counter"`
	CounterPadLength int    `gorm:"type:INT" json:"counter_pad_length"`
	CounterPadChar   string `gorm:"type:CHAR(1)" json:"counter_pad_char"`
	ExtraField       string `gorm:"type:VARCHAR(255)" json:"extra_field"`
}

func (RefFormat) TableName() string {
	return "ref_format"
}

func GenerateRefferalCode(ctx context.Context, DB *gorm.DB) string {
	var refFormat RefFormat
	err := DB.WithContext(ctx).Where("format_id = ?", "referral_code").First(&refFormat).Error
	exception.PanicIfError(err)
	refFormat.Counter += 1
	err = DB.WithContext(ctx).Where("format_id = ?", "referral_code").Updates(&refFormat).Error
	exception.PanicIfError(err)
	code := fmt.Sprintf("%v%0*d", refFormat.ExtraField, refFormat.CounterPadLength, refFormat.Counter)
	return code
}
