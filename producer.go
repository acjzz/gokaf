package gokaf

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

// Producer struct represents a message producer in the pubsub system.
type Producer struct {
	id     uuid.UUID
	ctx    context.Context
	cancel context.CancelFunc
	logger *slog.Logger
	wg     sync.WaitGroup // Add a WaitGroup for synchronization
	topic  *topic
}

func newProducer(topic *topic, logger *slog.Logger) *Producer {
	ctx, cancel := context.WithCancel(topic.ctx)

	p := Producer{uuid.New(), ctx, cancel, logger, sync.WaitGroup{}, topic}

	p.wg.Add(1) // Increment the WaitGroup counter
	return &p
}

// Stop gracefully stops the producer by canceling its context and waiting for associated tasks to complete.
func (p *Producer) Stop() {
	defer p.wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	// Shutdown. Cancel application context will kill all attached tasks.
	p.logger.Warn(fmt.Sprintf("Producer[%s] for topic %s close", p.id, p.topic.name))
	p.cancel()
}

// Publish sends a message to the associated topic through the producer.
func (p *Producer) Publish(message interface{}) error {
	return p.topic.publish(message)
}
