package request

import (
	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
)

type CreateTransactionRequest struct {
	Key   string           `json:"key"`   //UserID
	Value ValueTransaction `json:"value"` //Amount
}

type ValueTransaction struct {
	Amount      int64                      `json:"amount" validate:"required"`
	Sender      *string                    `json:"sender"`
	Date        string                     `json:"date"`
	Reciver     *string                    `json:"reciver"`
	UserID      string                     `json:"user_id" validate:"required"`
	Currency    string                     `json:"currency" validate:"required"`
	ProcessType string                     `json:"process_type" validate:"required"`
	Status      constant.StatusTransaction `json:"status" validate:"required"`
	PaymentType constant.PaymentAccount    `json:"payment_type" validate:"required"`
}

type ReqCreateLoan struct {
	Amount      int64  `json:"amount" validate:"required"`
	UserID      string `json:"user_id" validate:"required"`
	Tenor       int    `json:"tenor" validate:"required"`
	Currency    string `json:"currency" validate:"required"`
	ProcessType string `json:"process_type" validate:"required"`
}

type ReqListTransaction struct {
	BaseQuery
	UserID string                     `json:"user_id"`
	Status constant.StatusTransaction `json:"status"`
}

// export type Transaction = {
//   amount: number;
//   sender: string;
//   reciver: string;
//   currency: string;
//   date: string;
//   status: Status;
//   processType: ProcessType;
// };

// export type ProcessType = "withdraw" | "transfer";

// export type Status = "pending" | "completed" | "failed";
