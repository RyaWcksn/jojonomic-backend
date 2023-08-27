package message

import (
	"strings"

	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/pkgs/logger"
	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(cfg config.Config, log logger.ILogger) *kafka.Reader {
	brokers := strings.Split(cfg.KafkaAddr, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    cfg.KafkaTopic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
