package user

import (
	"database/sql"
	"time"

	"github.com/a-fandy/finan/api/v1/income"
	"gorm.io/gorm"
)

type User struct {
	Id           uint64          `gorm:"primaryKey" json:"id"`
	Username     string          `gorm:"type:VARCHAR(255)" json:"username"`
	Password     string          `gorm:"type:VARCHAR(255) NOT NULL" json:"password"`
	Email        string          `gorm:"type:VARCHAR(255) NOT NULL UNIQUE" json:"email"`
	Name         string          `gorm:"type:VARCHAR(255) NOT NULL" json:"name"`
	NoHp         string          `gorm:"type:VARCHAR(15) NOT NULL" json:"no_hp"`
	Role         string          `gorm:"type:VARCHAR(15) NOT NULL" json:"role"`
	RefferalCode string          `gorm:"type:VARCHAR(15) NULL" json:"refferal_code"`
	Status       bool            `gorm:"type:TINYINT NOT NULL" json:"status"`
	VerifyAt     sql.NullTime    `json:"verify_at"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"deleted_at"`
	Incomes      []income.Income `json:"incomes"`
}

// func (User) TableName() string {
// 	return "tb_user"
// }

type UserResponse struct {
	Id           uint64 `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	NoHp         string `json:"no_hp"`
	Status       bool   `json:"status"`
	RefferalCode string `json:"refferal_code"`
}

func UserEntityToResponse(user User) UserResponse {
	return UserResponse{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		Name:         user.Name,
		NoHp:         user.NoHp,
		RefferalCode: user.RefferalCode,
		Status:       user.Status,
	}
}
