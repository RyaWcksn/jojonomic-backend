package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Service  string
	LogLevel string

	Host string
	Port string

	KafkaAddr  string
	KafkaTopic string
}

func InitCfg() *Config {
	if os.Getenv("ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error while init config... %v", err)
		}
	}

	env := os.Getenv("ENV")
	service := os.Getenv("SERVICE")
	logLevel := os.Getenv("LOG_LEVEL")

	kafkaAddr := os.Getenv("KAFKA_ADDR")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return &Config{
		Env:        env,
		Service:    service,
		LogLevel:   logLevel,
		Host:       host,
		Port:       port,
		KafkaAddr:  kafkaAddr,
		KafkaTopic: kafkaTopic,
	}
}
