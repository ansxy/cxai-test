package custom_db

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"gorm.io/gorm"
)

func AutoMigrateDatabse(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Account{},
		&model.Transaction{},
		&model.Loan{},
		&model.LoanLog{},
	)
}
