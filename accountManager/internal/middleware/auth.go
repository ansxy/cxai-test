package middleware

import (
	"context"

	"github.com/ansxy/golang-boilerplate-gin/config"
	"github.com/ansxy/golang-boilerplate-gin/internal/response"
	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
	custom_error "github.com/ansxy/golang-boilerplate-gin/pkg/error"
	"github.com/ansxy/golang-boilerplate-gin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/supabase-community/supabase-go"
)

type Middleware struct {
	Config   *config.JWTConfig
	Supabase *supabase.Client
}

func NewMiddleware(config *config.JWTConfig, supabase *supabase.Client) *Middleware {
	return &Middleware{Config: config, Supabase: supabase}
}

func (m *Middleware) AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetTokenFromHeader(c.Request)
		if err != nil {
			response.Error(c.Writer, err)
			c.Abort()
			return
		}

		jwtClaim := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(token, jwtClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.Config.JWTSecret), nil
		})
		if err != nil {
			err = custom_error.SetCustomeError(&custom_error.ErrorContext{
				Code:    constant.DefaultUnauthorizedErrorCode,
				Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedErrorCode],
			})
			response.Error(c.Writer, err)
			c.Abort()
			return
		}

		userID, ok := jwtClaim["sub"].(string)
		if !ok {
			err = custom_error.SetCustomeError(&custom_error.ErrorContext{
				Code:    constant.DefaultUnauthorizedErrorCode,
				Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedErrorCode],
			})
			response.Error(c.Writer, err)
			c.Abort()
			return
		}

		// log.Println("userID", userID)
		// customClaims, ok := resJwt.Claims.(*custom_jwt.UserClaims)
		// if !ok || !resJwt.Valid {
		// 	err = custom_error.SetCustomeError(&custom_error.ErrorContext{
		// 		Code:    constant.DefaultUnauthorizedErrorCode,
		// 		Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedErrorCode],
		// 	})
		// 	response.Error(c.Writer, err)
		// 	c.Abort()
		// 	return
		// }

		ctx := context.WithValue(c.Request.Context(), constant.USERIDKEY, userID)
		// ctx = context.WithValue(ctx, constant.ROLEKEY, customClaims.Role)
		// c.Request = c.Request.WithContext(ctx)
		c.Request = c.Request.WithContext(ctx)
		// Continue to the next handler
		c.Next()
	}
}
