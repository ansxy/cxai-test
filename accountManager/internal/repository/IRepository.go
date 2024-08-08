package repository

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/ansxy/golang-boilerplate-gin/internal/request"
	"github.com/gin-gonic/gin"
)

type IFaceRepository interface {
	CreateUser(c *gin.Context, data *model.User) error

	GetUser(c *gin.Context, query ...interface{}) (*model.User, error)

	GetListAccount(c *gin.Context, userID string) ([]model.Account, error)
	GetListTransaction(c *gin.Context, req request.ReqListTransaction) ([]model.Transaction, int64, error)
	CreateAccount(c *gin.Context, data *model.Account) error
}
