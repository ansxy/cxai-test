package http_transport

import (
	"net/http"

	"github.com/ansxy/golang-boilerplate-gin/config"
	"github.com/ansxy/golang-boilerplate-gin/internal/middleware"
	v1 "github.com/ansxy/golang-boilerplate-gin/internal/transport/http/v1"
	"github.com/ansxy/golang-boilerplate-gin/internal/usecase"
	custom_gin "github.com/ansxy/golang-boilerplate-gin/pkg/gin"
	"github.com/sirupsen/logrus"
	"github.com/supabase-community/supabase-go"
)

func NewHttpTransport(cnf *config.Config, uc usecase.IUsecase, supa *supabase.Client) http.Handler {
	app := custom_gin.NewGinApp()
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	mw := middleware.Middleware{
		Supabase: supa,
		Config:   &cnf.JWTConfig,
	}
	app.Use(middleware.NewLogger(logger))

	v1.NewV1Transport(app.Group("/api"), uc, mw)

	return app
}
