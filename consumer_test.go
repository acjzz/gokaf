package gokaf

import (
	"context"
	"testing"
)

func Test_consumer(t *testing.T) {
	t.Run("Consumer", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		var channel chan internalMessage
		c := newConsumer(ctx, &channel, func(s string, i interface{}) {})
		c.run()
		cancel()
	})
}
