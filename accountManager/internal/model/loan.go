package model

import "gorm.io/gorm"

type Loan struct {
	LoanID       string  `json:"loan_id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID       string  `json:"user_id"`
	Amount       float64 `json:"amount"`
	InterestRate float64 `json:"interest_rate"`
	Tenor        int     `json:"tenor"`
	Currency     string  `json:"currency"`

	CreatedAt string          `json:"created_at" gorm:"timestamp;default:now()"`
	UpdatedAt string          `json:"updated_at" gorm:"timestamp;default:now()"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"timestamp"`

	LoanLogs []LoanLog `json:"loan_logs" gorm:"foreignKey:LoanID;references:LoanID"`
}

func (l *Loan) TableName() string {
	return "tr_loan"
}
