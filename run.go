package main

import (
	gofka "./src"
	"context"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t := gofka.NewTopic(ctx, "test")
	c := t.NewConsumer()
	c.Run()
	p := t.NewProducer()
	p.Publish(gofka.NewInternalMessage("Testing"))
}
