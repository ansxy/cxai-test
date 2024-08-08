package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	AccountID uuid.UUID      `json:"account_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID      `json:"user_id" gorm:"type:uuid"`
	Type      string         `json:"type"`
	Balance   float64        `json:"balance"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (a *Account) TableName() string {
	return "tr_account"
}
