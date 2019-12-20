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
	fmt.Printf("%s\n", getEngineToken(ctx))
	return &consumer{ctx, ch}
}

func (c *consumer) run() {
	go func() {
		fmt.Println("Consumer start")
		for {
			select {
			case <- c.ctx.Done():
				fmt.Println("Consumer stop")
				return
			case m, ok := <-*c.channel:
				if !ok {
					fmt.Println("Topic closed")
					break
				} else {
					m.consume()
				}
			}
		}
	}()
}
