package user

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/middleware"
	"github.com/ansxy/golang-boilerplate-gin/internal/usecase"
	custom_validator "github.com/ansxy/golang-boilerplate-gin/pkg/validator"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	v  custom_validator.Validator
	uc usecase.IUsecase
	mw middleware.Middleware
}

func NewUserHandler(router *gin.RouterGroup, v custom_validator.Validator, uc usecase.IUsecase, mw middleware.Middleware) *gin.RouterGroup {

	handler := &UserHandler{
		v:  v,
		uc: uc,
		mw: mw,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
	}

	account := router.Group("/account")
	{
		account.Use(mw.AuthenticateUser())
		account.GET("/", handler.GetAccount)
	}

	transaction := router.Group("/transaction")
	{
		transaction.Use(mw.AuthenticateUser())
		transaction.POST("/transfer", handler.SendBalance)
		transaction.POST("/withdraw", handler.WithdrawBalance)
		transaction.POST("/loan", handler.TakeAloan)
		transaction.GET("/", handler.GetListTransaction)
	}

	return router
}
