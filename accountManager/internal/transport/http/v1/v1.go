package v1

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/handler/public"
	"github.com/ansxy/golang-boilerplate-gin/internal/handler/user"
	"github.com/ansxy/golang-boilerplate-gin/internal/middleware"
	"github.com/ansxy/golang-boilerplate-gin/internal/usecase"
	custom_validator "github.com/ansxy/golang-boilerplate-gin/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewV1Transport(route *gin.RouterGroup, uc usecase.IUsecase, mw middleware.Middleware) *gin.RouterGroup {
	validator := custom_validator.NewValidator(validator.New())
	routeGroup := route.Group("/v1")

	public.NewPublicHandler(routeGroup, validator)
	user.NewUserHandler(routeGroup, validator, uc, mw)
	return route
}
