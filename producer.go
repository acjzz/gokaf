package gokaf

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

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

func (p *Producer) Close() {
	defer p.wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	// Shutdown. Cancel application context will kill all attached tasks.
	p.logger.Warn(fmt.Sprintf("Producer[%s] for topic %s close", p.id, p.topic.name))
	p.cancel()
}

func (p *Producer) Publish(message interface{}) error {
	return p.topic.publish(message)
}
