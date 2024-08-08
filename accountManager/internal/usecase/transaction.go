package usecase

import (
	"encoding/json"

	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/ansxy/golang-boilerplate-gin/internal/request"
	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// CreateTransaction implements IUsecase.
func (u *Usecase) CreateTransaction(c *gin.Context, data request.ValueTransaction) (map[string]string, error) {
	valueBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msg := &kafka.Message{
		Key:   []byte(data.UserID),
		Value: valueBytes,
	}

	err = u.Svc.SendMessage(c, *msg)

	if err != nil {
		return nil, err
	}

	res := map[string]string{
		"message": "Transaction is being processed",
		"status":  string(constant.Pending),
	}

	return res, nil
}

// GetListTransaction implements IUsecase.
func (u *Usecase) GetListTransaction(c *gin.Context, params request.ReqListTransaction) ([]model.Transaction, int64, error) {
	return u.Repo.GetListTransaction(c, params)
}
