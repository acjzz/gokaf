package src

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
)

type producer struct {
	ctx     context.Context
	channel *chan internalMessage
	logger  *logrus.Entry
}

func newProducer(ctx context.Context, ch *chan internalMessage) *producer {
	return &producer{ctx, ch, logrus.WithFields(getLogFields(ctx))}
}

func (p *producer) publish(message internalMessage) error {
	select {
	case <-p.ctx.Done():
		p.logger.Warn("closed")
		return fmt.Errorf("Topic closed")
	default:
		*p.channel <- message
		return nil
	}
}
