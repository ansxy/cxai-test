package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;"`
	Email  string    `json:"email"`

	CreatedAt string          `json:"created_at" gorm:"timestamp;default:now()"`
	UpdatedAt string          `json:"updated_at" gorm:"timestamp;default:now()"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"timestamp"`

	Accounts []*Account `json:"accounts" gorm:"foreignKey:UserID;references:UserID"`
}

func (u *User) TableName() string {
	return "tr_user"
}
