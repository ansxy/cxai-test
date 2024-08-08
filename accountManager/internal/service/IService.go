package service

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type IService interface {
	SendMessage(ctx *gin.Context, data kafka.Message) error
}
