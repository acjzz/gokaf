package src

import "context"

type producer struct {
	ctx		context.Context
	channel *chan messageInterface
}

func newProducer(ctx context.Context, ch *chan messageInterface) *producer {
	return &producer{ctx, ch}
}

func (p *producer) publish(message messageInterface) {
	*p.channel <- message
}
