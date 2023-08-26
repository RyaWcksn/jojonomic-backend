package broker

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/logger"
	"github.com/segmentio/kafka-go"
)

type IBroker interface {
	Consume(ctx context.Context) <-chan []byte
}

type BrokerImpl struct {
	kafkaConn *kafka.Reader
	l         logger.ILogger
}

var _ IBroker = (*BrokerImpl)(nil)

func NewBrokerImpl(k *kafka.Reader, l logger.ILogger) *BrokerImpl {
	return &BrokerImpl{
		kafkaConn: k,
		l:         l,
	}
}
