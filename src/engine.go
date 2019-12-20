package src

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"strings"
)

type GofkaEngine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	logger    *logrus.Entry
	topics    map[string]*Topic
}

func NewGofkaEngine(name string) *GofkaEngine {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = setEngineKey(ctx, name)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	ge := GofkaEngine{
		ctx,
		cancel,
		logrus.WithFields(getLogFields(ctx)),
		map[string]*Topic{},
	}

	go func(ge *GofkaEngine) {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		// Shutdown. Cancel application context will kill all attached tasks.
		ge.logger.Warn("shutting down")
		ge.Stop()
	}(&ge)

	return &ge
}

func (ge *GofkaEngine) Stop() {
	ge.ctxCancel()
}

func (ge *GofkaEngine) AddTopic(name string, numConsumers ...int ) {
	name = strings.ToLower(name)
	if _, ok := ge.topics[name]; !ok {
		ctx := setTopicKey(ge.ctx, name)
		ge.topics[name] = NewTopic(ctx, name)
		if len(numConsumers) > 0 {
			ge.topics[name].AddConsumers(numConsumers[0])
		}
		ge.topics[name].Run()
	} else {
		ge.logger.Warn("topic already exists")
	}
}

func (ge *GofkaEngine) Publish(name string, obj interface{}) error {
	name = strings.ToLower(name)
	if _, ok := ge.topics[name]; ok {
		return ge.topics[name].Publish(newInternalMessage(obj))
	} else {
		ge.logger.Error("topic does not exist")
		return fmt.Errorf("Topic %s does not exists\n", name)
	}
}
