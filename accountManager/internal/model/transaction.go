package model

import "github.com/google/uuid"

type Transaction struct {
	TransactionID uuid.UUID `json:"transaction_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID        uuid.UUID `json:"user_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Sender        string    `json:"sender"`
	Receiver      string    `json:"receiver"`

	Type      string `json:"type"`
	Status    string `json:"status" gorm:"enum('pending', 'success', 'failed')"`
	CreatedAt string `json:"created_at" gorm:"timestamp;default:now()"`
	UpdatedAt string `json:"updated_at" gorm:"timestamp;default:now()"`
}

func (t *Transaction) TableName() string {
	return "tr_transaction"
}
