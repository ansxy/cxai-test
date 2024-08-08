package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/ansxy/golang-boilerplate-gin/internal/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

// CreateAccount implements IUsecase.
func (u *Usecase) CreateAccount(c *gin.Context, data request.ReqCreateAccount) error {
	// if data.Type == string(constant.Debit) {
	// 	//Create Scheduler for every day on month data.schedule
	// 	panic("Not implemented")
	// }
	// if data.Type == string(constant.Credit) {
	// 	panic("Not implemented")
	// }

	userID, err := uuid.Parse(data.UserID)
	if err != nil {
		return err
	}

	account := &model.Account{
		UserID:  userID,
		Type:    data.Type,
		Balance: data.Balance,
	}

	return u.Repo.CreateAccount(c, account)
}

// TakeLoan implements IUsecase.
func (u *Usecase) TakeLoan(c *gin.Context, data request.ReqCreateLoan) error {
	valueBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := &kafka.Message{
		Key:   []byte(fmt.Sprintf("[LOAN] - %s", data.UserID)),
		Value: valueBytes,
	}

	err = u.Svc.SendMessage(c, *msg)
	if err != nil {
		return err
	}

	return nil
}

// GetAccount implements IUsecase.
func (u *Usecase) GetAccount(c *gin.Context, userID string) (*model.User, error) {
	return u.Repo.GetUser(c, "user_id = ?", userID)
}
