package gokaf

import (
	"context"
	"fmt"
)

type producer struct {
	ctx     context.Context
	channel *chan internalMessage
	logger  logWrapper
}

func newProducer(ctx context.Context, ch *chan internalMessage) *producer {
	pctx := setProducerKey(ctx)
	return &producer{pctx, ch, NewLogger(pctx)}
}

func (p *producer) publish(message internalMessage) error {
	select {
	case <-p.ctx.Done():
		p.logger.Warn("closed")
		return fmt.Errorf("topic closed")
	default:
		*p.channel <- message
		return nil
	}
}
