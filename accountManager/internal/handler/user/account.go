package user

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/request"
	"github.com/ansxy/golang-boilerplate-gin/internal/response"
	"github.com/ansxy/golang-boilerplate-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetAccount(c *gin.Context) {
	userID, err := utils.GetUserIDFromCtx(c.Request.Context())
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	user, err := h.uc.GetAccount(c, userID)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	response.Success(c.Writer, user)

}

func (h *UserHandler) TakeAloan(c *gin.Context) {
	var req request.ReqCreateLoan
	userId, err := utils.GetUserIDFromCtx(c.Request.Context())
	if err != nil {
		response.Error(c.Writer, err)
		return
	}
	req.ProcessType = "loan"
	req.UserID = userId
	err = h.v.ValidateStruct(c.Request, &req)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	err = h.uc.TakeLoan(c, req)

	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	response.Success(c.Writer, "Loan success")
}

func (h *UserHandler) AddAccount(c *gin.Context) {
	var req request.ReqCreateAccount
	userId, err := utils.GetUserIDFromCtx(c.Request.Context())
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	req.UserID = userId
	err = h.v.ValidateStruct(c.Request, &req)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	err = h.uc.CreateAccount(c, req)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	response.Success(c.Writer, nil)
}
