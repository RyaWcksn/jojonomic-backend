package config

import (
	"log"
	"os"
	"strconv"

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

	SQLAddr       string
	SQLPort       int
	SQLUser       string
	SQLPassword   string
	SQLDatabase   string
	MaxConnection int
	MaxOpen       int
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

	dbHost := os.Getenv("SQL_ADDR")
	dbPort, _ := strconv.Atoi(os.Getenv("SQL_PORT"))
	dbUser := os.Getenv("SQL_USER")
	dbPass := os.Getenv("SQL_PASS")
	db := os.Getenv("SQL_DB")
	maxConn, _ := strconv.Atoi(os.Getenv("SQL_MAX_CONN"))
	maxOpen, _ := strconv.Atoi(os.Getenv("SQL_MAX_OPEN"))

	return &Config{
		Env:           env,
		Service:       service,
		LogLevel:      logLevel,
		Host:          host,
		Port:          port,
		KafkaAddr:     kafkaAddr,
		KafkaTopic:    kafkaTopic,
		SQLAddr:       dbHost,
		SQLPort:       dbPort,
		SQLUser:       dbUser,
		SQLPassword:   dbPass,
		SQLDatabase:   db,
		MaxConnection: maxConn,
		MaxOpen:       maxOpen,
	}
}
