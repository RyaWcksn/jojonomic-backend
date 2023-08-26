package message

import (
	"context"
	"fmt"

	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/logger"
	"github.com/segmentio/kafka-go"
)

func NewKafkaProducer(cfg config.Config, log logger.ILogger) *kafka.Conn {
	broker, err := kafka.DialLeader(context.Background(), "tcp", cfg.KafkaAddr, cfg.KafkaTopic, 0)
	if err != nil {
		fmt.Println("ERROR := ", err.Error())
		if err == kafka.InvalidTopic {

		}

		log.Fatal(err.Error())
	}

	return broker
}

func createTopic(cfg config.Config, log logger.ILogger) {

}
