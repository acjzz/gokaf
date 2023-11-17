package gokaf

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

type Consumer struct {
	id      uuid.UUID
	ctx     context.Context
	cancel  context.CancelFunc
	logger  *slog.Logger
	wg      sync.WaitGroup // Add a WaitGroup for synchronization
	topic   *topic
	handler func(interface{})
}

func newConsumer(topic *topic, logger *slog.Logger, handler func(interface{})) *Consumer {
	ctx, cancel := context.WithCancel(topic.ctx)

	c := Consumer{uuid.New(), ctx, cancel, logger, sync.WaitGroup{}, topic, handler}

	c.wg.Add(1) // Increment the WaitGroup counter
	return &c
}

func (c *Consumer) Close() {
	defer c.wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	// Shutdown. Cancel application context will kill all attached tasks.
	c.logger.Warn(fmt.Sprintf("Consumer[%s] for topic %s close", c.id, c.topic.name))
	c.cancel()
}

func (c *Consumer) Run() {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			case msg := <-c.topic.channel.ch:
				c.handler(msg)
			}
		}
	}()
}
