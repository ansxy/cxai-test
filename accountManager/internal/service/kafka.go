package service

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// SendMessage implements IService.
func (s *Service) SendMessage(ctx *gin.Context, data kafka.Message) error {
	if err := s.Kafka.Writer.WriteMessages(ctx, data); err != nil {
		return err
	}

	return nil
}
