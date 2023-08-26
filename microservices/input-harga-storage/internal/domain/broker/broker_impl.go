package broker

import (
	"context"
)

// Consume implements IBroker.
func (k *BrokerImpl) Consume(ctx context.Context) <-chan []byte {
	messages := make(chan []byte)

	go func() {
		defer close(messages)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				m, err := k.kafkaConn.FetchMessage(ctx)
				if err != nil {
					k.l.Errorf("Error while reading message := %v", err)
					return
				}
				messages <- m.Value
			}
		}
	}()
	return messages
}
