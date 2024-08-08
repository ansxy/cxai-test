package user

import (
	"log"

	"github.com/ansxy/golang-boilerplate-gin/internal/request"
	"github.com/ansxy/golang-boilerplate-gin/internal/response"
	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
	"github.com/ansxy/golang-boilerplate-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) SendBalance(c *gin.Context) {
	var req request.ValueTransaction
	userId, err := utils.GetUserIDFromCtx(c.Request.Context())
	if err != nil {
		response.Error(c.Writer, err)
		return
	}
	req.Date = utils.DateTimeNow()
	req.UserID = userId
	req.Status = constant.Pending
	req.ProcessType = "transfer"
	req.Sender = &userId

	err = h.v.ValidateStruct(c.Request, &req)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	res, err := h.uc.CreateTransaction(c, req)
	if err != nil {
		log.Println("err", err)
		response.Error(c.Writer, err)
	}

	response.Success(c.Writer, res)

}

func (h *UserHandler) WithdrawBalance(c *gin.Context) {
	var req request.ValueTransaction

	userId, err := utils.GetUserIDFromCtx(c.Request.Context())
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	req.Date = utils.DateTimeNow()
	req.UserID = userId
	req.Status = constant.Pending
	req.ProcessType = "withdraw"

	err = h.v.ValidateStruct(c.Request, &req)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	res, err := h.uc.CreateTransaction(c, req)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	response.Success(c.Writer, res)

}

func (h *UserHandler) GetListTransaction(c *gin.Context) {
	var req request.ReqListTransaction

	req.BaseQuery = request.NewPaginationQuery(c.Request)
	req.Status = constant.StatusTransaction(c.Param("status"))
	userId, err := utils.GetUserIDFromCtx(c.Request.Context())
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	req.UserID = userId

	res, cnt, err := h.uc.GetListTransaction(c, req)

	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	response.Pagination(c.Writer, res, req.Page, req.PerPage, cnt)

}
