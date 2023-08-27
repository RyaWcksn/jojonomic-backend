package broker

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/logger"
	"github.com/segmentio/kafka-go"
)

//go:generate mockgen -source broker.go -destination broker_mock.go -package broker
type IBroker interface {
	Publish(ctx context.Context, message *BrokerMessage) error
}

type BrokerImpl struct {
	kafkaConn *kafka.Conn
	log       logger.ILogger
}

var _ IBroker = (*BrokerImpl)(nil)

func NewMessageBroker(broker *kafka.Conn, l logger.ILogger) *BrokerImpl {
	return &BrokerImpl{
		kafkaConn: broker,
		log:       l,
	}
}
