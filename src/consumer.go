package src

import (
	"context"
	"fmt"
)

type consumer struct {
	ctx		context.Context
	channel *chan messageInterface
}

func newConsumer(ctx context.Context, ch *chan messageInterface) *consumer {
	return &consumer{ctx, ch}
}

func (c *consumer) Run() {
	go func() {
		for {
			select {
			case <- c.ctx.Done():
				return
			case m, ok := <-*c.channel:
				if !ok {
					fmt.Println("topic closed")
					break
				} else {
					m.consume()
				}
			}
		}
	}()
}
