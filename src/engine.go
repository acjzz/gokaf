package src

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

type GofkaEngine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	topics    map[string]*Topic
}

func NewGofkaEngine(name string) *GofkaEngine {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = setEngineKey(ctx, name)
	ge := GofkaEngine{ ctx,cancel, map[string]*Topic{}, }

	go func(ge *GofkaEngine) {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		// Shutdown. Cancel application context will kill all attached tasks.
		fmt.Println("Shutting down gofka engine")
		ge.Stop()
	}(&ge)

	return &ge
}

func (ge *GofkaEngine) Stop() {
	ge.ctxCancel()
}

func (ge *GofkaEngine) AddTopic(name string) {
	name = strings.ToLower(name)
	if _, ok := ge.topics[name]; !ok {
		ctx := setTopicKey(ge.ctx, name)
		ge.topics[name] = NewTopic(ctx, name)
		ge.topics[name].Run()
	} else {
		fmt.Printf("Topic %s already exists\n", name)
	}
}

func (ge *GofkaEngine) Publish(name string, message messageInterface) error {
	name = strings.ToLower(name)
	if _, ok := ge.topics[name]; ok {
		return ge.topics[name].Publish(message)
	} else {
		return fmt.Errorf("Topic %s does not exists\n", name)
	}
}
