package gokaf

import (
	"context"
	"github.com/sirupsen/logrus"
)

type consumer struct {
	ctx     context.Context
	channel *chan internalMessage
	logger  *logrus.Entry
	handler func(interface{})
}

func newConsumer(ctx context.Context, ch *chan internalMessage, handler func(interface{})) *consumer {
	return &consumer{ctx, ch, logrus.WithFields(getLogFields(ctx)), handler, }
}

func (c *consumer) run() {
	go func() {
		c.logger.Debug("Start")
		for {
			select {
			case <- c.ctx.Done():
				c.logger.Debug("Stop")
				return
			case m, ok := <-*c.channel:
				if !ok {
					c.logger.Warn("Closed")
					break
				} else {
					c.logger.Tracef("Consume => %s", m.value)
					c.handler(m.value)
				}
			}
		}
	}()
}
