package gokaf

import (
	"context"
)

type consumer struct {
	ctx     context.Context
	channel *chan internalMessage
	logger  logWrapper
	handler func(string, interface{})
}

func newConsumer(ctx context.Context, ch *chan internalMessage, handler func(string, interface{})) *consumer {
	return &consumer{ctx, ch, NewLogger(ctx), handler}
}

func (c *consumer) run() {
	go func() {
		c.logger.Debug("Start")
		for {
			select {
			case <-c.ctx.Done():
				c.logger.Debug("stop")
				return
			case m, ok := <-*c.channel:
				if !ok {
					c.logger.Warn("Closed")
					break
				} else {
					c.logger.Tracef("Consume => %s", m.value)
					c.handler(getTopicKey(c.ctx), m.value)
				}
			}
		}
	}()
}
