package gokaf

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test_consumer(t *testing.T) {
	t.Run("Consumer", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ctx = setLogLevelKey(ctx, logrus.InfoLevel)

		var channel chan internalMessage
		c := newConsumer(ctx, &channel, func(s string, i interface{}) interface{}{ return nil })
		c.run()
		cancel()
	})
}
