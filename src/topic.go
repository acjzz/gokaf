package src

import "context"

type Topic struct {
	ctx		context.Context
	name	string
	channel chan messageInterface
}

func NewTopic(ctx context.Context, name string) *Topic {
	channelTopic := make(chan messageInterface)
	return &Topic{
		ctx,
		name,
		channelTopic,
	}
}

func (t *Topic) NewConsumer()  *consumer {
	return newConsumer(t.ctx, &t.channel)
}

func (t *Topic) NewProducer()  *producer {
	return newProducer(t.ctx, &t.channel)
}
