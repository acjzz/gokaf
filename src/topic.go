package src

import (
	"context"
)

type Topic struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	name      string
	channel   chan messageInterface
	consumers []*consumer
	producer  *producer
}

func NewTopic(ctx context.Context, name string) *Topic {
	channelTopic := make(chan messageInterface)
	ctx, cancel := context.WithCancel(ctx)
	pctx := setProducerKey(ctx)

	return &Topic{
		ctx,
		cancel,
		name,
		channelTopic,
		[]*consumer{},
		newProducer(pctx, &channelTopic),
	}
}

func (t *Topic) Stop() {
	t.ctxCancel()
}

func (t *Topic) AddConsumer() {
	ctx := setConsumerKey(t.ctx, len(t.consumers))
	t.consumers = append(t.consumers, newConsumer(ctx, &t.channel))
}

func (t *Topic) AddConsumers(num int) {
	for i := 0; i < num; i += 1 {
		t.AddConsumer()
	}
}

func (t *Topic) Publish(message messageInterface) error {
	return t.producer.publish(message)
}

func (t *Topic) Run() {
	if len(t.consumers) == 0 { t.AddConsumer() }
	for _, c := range t.consumers {
		c.run()
	}
}
