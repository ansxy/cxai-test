package model

import "time"

type LoanLog struct {
	LoanLogID   string    `json:"loan_log_id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	LoanID      string    `json:"loan_id"`
	TotalAmount float64   `json:"total_amount"`
	IsPaid      bool      `json:"is_paid"`
	Duedate     time.Time `json:"duedate"`

	CreatedAt string `json:"created_at" gorm:"timestamp;default:now()"`
	UpdatedAt string `json:"updated_at" gorm:"timestamp;default:now()"`
}

func (l *LoanLog) TableName() string {
	return "tr_loan_log"
}
