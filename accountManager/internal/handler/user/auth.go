package user

import (
	"log"

	"github.com/ansxy/golang-boilerplate-gin/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/gotrue-go/types"
)

func (h *UserHandler) Register(c *gin.Context) {
	var registerReq types.SignupRequest

	err := h.v.ValidateStruct(c.Request, &registerReq)

	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	err = h.uc.Register(c, registerReq)

	if err != nil {
		log.Println(err)
		response.Error(c.Writer, err)
		return
	}

	response.Success(c.Writer, nil)

}

func (h *UserHandler) Login(c *gin.Context) {
	var Login types.SignupRequest

	err := h.v.ValidateStruct(c.Request, &Login)

	if err != nil {
		response.Error(c.Writer, err)
	}

	token, err := h.uc.Login(c, Login)

	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	response.Success(c.Writer, token)
}
