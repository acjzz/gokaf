package src

import (
	"context"
	"fmt"
)

type producer struct {
	ctx		context.Context
	channel *chan messageInterface
}

func newProducer(ctx context.Context, ch *chan messageInterface) *producer {
	return &producer{ctx, ch}
}

func (p *producer) publish(message messageInterface) error {
	select {
	case <-p.ctx.Done():
		return fmt.Errorf("Topic closed")
	default:
		*p.channel <- message
		return nil
	}
}
