package broker

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	rr "github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/errors"
	"github.com/segmentio/kafka-go"
)

// Publish implements IBroker.
func (b *BrokerImpl) Publish(ctx context.Context, message *BrokerMessage) error {

	b.kafkaConn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	bytePayload, err := json.Marshal(&message)
	if err != nil {
		b.log.Errorf("Error while marshaling payload := %v", err)
		return rr.GetError(message.ReffId, err)
	}

	msg := kafka.Message{
		Value: []byte(bytePayload),
	}

	_, err = b.kafkaConn.WriteMessages(msg)
	if err != nil {
		b.log.Errorf("Error while marshaling payload := %v", err)
		return rr.GetError(message.ReffId, errors.New("Kafka not ready"))
	}

	return nil

}
