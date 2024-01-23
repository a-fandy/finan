package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           uint64         `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"type:VARCHAR(255)" json:"username"`
	Password     string         `gorm:"type:VARCHAR(255) NOT NULL" json:"password"`
	Email        string         `gorm:"type:VARCHAR(255) NOT NULL UNIQUE" json:"email"`
	Name         string         `gorm:"type:VARCHAR(255) NOT NULL" json:"name"`
	NoHp         string         `gorm:"type:VARCHAR(15) NOT NULL" json:"no_hp"`
	Role         string         `gorm:"type:VARCHAR(15) NOT NULL" json:"role"`
	RefferalCode string         `gorm:"type:VARCHAR(15) NULL" json:"refferal_code"`
	Status       bool           `gorm:"type:TINYINT NOT NULL" json:"status"`
	VerifyAt     sql.NullTime   `json:"verify_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

// func (User) TableName() string {
// 	return "tb_user"
// }
