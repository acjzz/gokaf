package src

import "context"

type producer struct {
	ctx		context.Context
	channel *chan messageInterface
}

func (p *producer) Publish(message messageInterface) {
	*p.channel <- message
}

func newProducer(ctx context.Context, ch *chan messageInterface) *producer {
	return &producer{ctx, ch}
}
