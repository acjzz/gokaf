package main

import (
	gofka "./src"
	"context"
	"os"
	"os/signal"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t := gofka.NewTopic(ctx, "test")
	c := t.NewConsumer()
	c.Run()
	p := t.NewProducer()
	p.Publish(gofka.NewInternalMessage("Testing"))
	// Wait for SIGINT.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	// Shutdown. Cancel application context will kill all attached tasks.
	cancel()
}
