package src

import (
	"context"
	"github.com/Sirupsen/logrus"
)

type consumer struct {
	ctx		context.Context
	channel *chan internalMessage
	logger    *logrus.Entry
}

func newConsumer(ctx context.Context, ch *chan internalMessage) *consumer {
	return &consumer{ctx, ch, logrus.WithFields(getLogFields(ctx))}
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
					c.logger.Infof("Consume => %s", m.value)
					m.consume()
				}
			}
		}
	}()
}
