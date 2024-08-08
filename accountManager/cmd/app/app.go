package app

import (
	"github.com/ansxy/golang-boilerplate-gin/config"
	"github.com/ansxy/golang-boilerplate-gin/internal/repository"
	"github.com/ansxy/golang-boilerplate-gin/internal/service"
	http_transport "github.com/ansxy/golang-boilerplate-gin/internal/transport/http"
	"github.com/ansxy/golang-boilerplate-gin/internal/usecase"
	custom_db "github.com/ansxy/golang-boilerplate-gin/pkg/database"
	custom_http "github.com/ansxy/golang-boilerplate-gin/pkg/http"
	custome_kafka "github.com/ansxy/golang-boilerplate-gin/pkg/kafka"
	custom_postgres "github.com/ansxy/golang-boilerplate-gin/pkg/postgres"
	"github.com/ansxy/golang-boilerplate-gin/pkg/supa"
)

func Exec() (err error) {
	cnf := config.NewConfig()

	supa, err := supa.NewSupaClient(&cnf.Supabase)
	if err != nil {
		return err
	}

	db, err := custom_postgres.NewPostgresClient(&cnf.Postgres)
	if err != nil {
		return err
	}

	err = custom_db.AutoMigrateDatabse(db)
	if err != nil {
		return err
	}

	kafka, err := custome_kafka.InitKafka(&cnf.KafkaConfig)
	if err != nil {
		return err
	}

	svc := service.NewService(&service.Service{
		Kafka: kafka,
	})

	repo := repository.NewRepository(db)

	uc := usecase.NewUsecase(&usecase.Usecase{
		SupaBase: supa,
		Repo:     repo,
		Svc:      svc,
	})

	httpTransport := http_transport.NewHttpTransport(cnf, uc, supa)
	err = custom_http.NewHttpServer(cnf, httpTransport)
	if err != nil {
		return err
	}

	return
}
