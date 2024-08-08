package usecase

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/ansxy/golang-boilerplate-gin/internal/request"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/gotrue-go/types"
)

type IUsecase interface {
	Register(c *gin.Context, data types.SignupRequest) error
	Login(c *gin.Context, data types.SignupRequest) (*types.TokenResponse, error)

	//Account
	CreateAccount(c *gin.Context, data request.ReqCreateAccount) error
	GetAccount(c *gin.Context, userID string) (*model.User, error)
	TakeLoan(c *gin.Context, data request.ReqCreateLoan) error

	//Transaction
	CreateTransaction(c *gin.Context, data request.ValueTransaction) (map[string]string, error)
	GetListTransaction(c *gin.Context, params request.ReqListTransaction) ([]model.Transaction, int64, error)
}
