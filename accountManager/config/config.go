package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App         App
	Postgres    PostgresConfig
	Supabase    SupabaseConfig
	JWTConfig   JWTConfig
	KafkaConfig KafkaConfig
}

type JWTConfig struct {
	JWTSecret string
}

type SupabaseConfig struct {
	ApiKey     string
	ProjectRef string
}

type KafkaConfigWriter struct {
	Brokers []string
	Topic   string
}

type KafkaConfigReader struct {
	Brokers   []string
	Topic     string
	Partition int
}

type KafkaConfig struct {
	KafkaConfigReader KafkaConfigReader
	KafkaConfigWriter KafkaConfigWriter
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	URI      string
}

type App struct {
	Port string
	Host string
	Name string
}

func NewConfig() *Config {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed-Read-Envrionment-File \n%v\n", err.Error())
	}

	v := viper.GetViper()

	viper.AutomaticEnv()

	return &Config{
		App: App{
			Port: v.GetString("APP_PORT"),
			Host: v.GetString("APP_HOST"),
			Name: v.GetString("APP_NAME"),
		},
		Postgres: PostgresConfig{
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetString("DB_PORT"),
			Database: v.GetString("DB_NAME"),
			User:     v.GetString("DB_USER"),
			Password: v.GetString("DB_PASSWORD"),
			URI:      fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", v.GetString("DB_USER"), v.GetString("DB_PASSWORD"), v.GetString("DB_HOST"), v.GetString("DB_PORT"), v.GetString("DB_NAME")),
		},
		Supabase: SupabaseConfig{
			ApiKey:     v.GetString("SUPA_KEY"),
			ProjectRef: v.GetString("SUPA_URL"),
		},
		JWTConfig: JWTConfig{
			JWTSecret: v.GetString("JWT_SECRET"),
		},
		KafkaConfig: KafkaConfig{
			KafkaConfigReader: KafkaConfigReader{},
			KafkaConfigWriter: KafkaConfigWriter{
				Brokers: []string{v.GetString("KAFKA_BROKER")},
				Topic:   v.GetString("KAFKA_TOPIC"),
			},
		},
	}
}
