package src

import (
	"context"
	"github.com/Sirupsen/logrus"
)

type Topic struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	logger    *logrus.Entry
	name      string
	channel   chan internalMessage
	consumers []*consumer
	producer  *producer
	handler   func(interface{})
}

func NewTopic(ctx context.Context, name string, handler func(interface{})) *Topic {
	channelTopic := make(chan internalMessage)
	ctx, cancel := context.WithCancel(ctx)
	pctx := setTopicKey(ctx, name)
	logger := logrus.WithFields(getLogFields(ctx))
	logger.Info("create")
	return &Topic{
		ctx,
		cancel,
		logger,
		name,
		channelTopic,
		[]*consumer{},
		newProducer(pctx, &channelTopic),
		handler,
	}
}

func (t *Topic) Stop() {
	t.logger.Warn("stop")
	t.ctxCancel()
}

func (t *Topic) AddConsumer() {
	ctx := setConsumerKey(t.ctx, len(t.consumers))
	t.consumers = append(t.consumers, newConsumer(ctx, &t.channel, t.handler))
}

func (t *Topic) AddConsumers(num int) {
	for i := 0; i < num; i += 1 {
		t.AddConsumer()
	}
}

func (t *Topic) Publish(message internalMessage) error {
	return t.producer.publish(message)
}

func (t *Topic) Run() {
	if len(t.consumers) == 0 { t.AddConsumer() }
	for _, c := range t.consumers {
		c.run()
	}
}
