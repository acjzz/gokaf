package gokaf

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Topic struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	logger    *logrus.Entry
	name      string
	channel   chan internalMessage
	consumers []*consumer
	producer  *producer
	handler   func(string, interface{})
}

func NewTopic(ctx context.Context, name string, handler func(string, interface{}), numConsumers ...int) *Topic {
	var channelTopic chan internalMessage
	if len(numConsumers) > 0 {
		channelTopic = make(chan internalMessage, numConsumers[0])
	} else {
		channelTopic = make(chan internalMessage)
	}
	ctx, cancel := context.WithCancel(ctx)
	pctx := setTopicKey(ctx, name)
	logger := logrus.WithFields(getLogFields(ctx))
	logger.Info("create")
	t := &Topic{
		ctx,
		cancel,
		logger,
		name,
		channelTopic,
		[]*consumer{},
		newProducer(pctx, &channelTopic),
		handler,
	}
	if len(numConsumers) > 0 {
		t.addConsumers(numConsumers[0])
	} else {
		t.addConsumer()
	}
	return t
}

func (t *Topic) Stop() {
	t.logger.Warn("stop")
	t.ctxCancel()
}

func (t *Topic) addConsumer() {
	ctx := setConsumerKey(t.ctx, len(t.consumers))
	t.consumers = append(t.consumers, newConsumer(ctx, &t.channel, t.handler))
}

func (t *Topic) addConsumers(num int) {
	for i := 0; i < num; i += 1 {
		t.addConsumer()
	}
}

func (t *Topic) Publish(message internalMessage) error {
	return t.producer.publish(message)
}

func (t *Topic) Run() {
	for _, c := range t.consumers {
		c.run()
	}
}
