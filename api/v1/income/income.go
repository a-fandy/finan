package income

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	Id        uint64         `gorm:"primaryKey" json:"id"`
	UserID    uint64         `json:"user_id"`
	Name      string         `gorm:"type:VARCHAR(255)" json:"name"`
	Amount    float64        `gorm:"type:DECIMAL(14,2)" json:"amount"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type IncomeRequest struct {
	Name   string  `json:"name" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type IncomeResponse struct {
	// Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func IncomeRequestToEntity(incomeRequest IncomeRequest) Income {
	var income Income
	income.Name = incomeRequest.Name
	income.Amount = incomeRequest.Amount
	return income
}

func IncomeEntityToResponse(income Income) IncomeResponse {
	return IncomeResponse{
		Name:      income.Name,
		Amount:    income.Amount,
		CreatedAt: income.CreatedAt,
	}
}
