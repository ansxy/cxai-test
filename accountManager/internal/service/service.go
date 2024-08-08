package service

import custome_kafka "github.com/ansxy/golang-boilerplate-gin/pkg/kafka"

type Service struct {
	Kafka *custome_kafka.KafaClient
}

func NewService(s *Service) IService {
	return &Service{
		Kafka: s.Kafka,
	}
}
