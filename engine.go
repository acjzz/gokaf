package gokaf

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
)

type Engine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	logger    LogWrapper
	topics    map[string]*Topic
}

func NewEngine(name string, logLevel logrus.Level) *Engine {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = setEngineKey(ctx, name)
	ctx = setLogLevelKeyInCtx(ctx, logLevel)
	ge := Engine{
		ctx,
		cancel,
		NewLogrusLogger(ctx, getLogFields),
		map[string]*Topic{},
	}

	go func(ge *Engine) {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		// Shutdown. Cancel application context will kill all attached tasks.
		ge.logger.Warn("shutting down")
		ge.Stop()
	}(&ge)

	return &ge
}

func (ge *Engine) Stop() {
	ge.ctxCancel()
}

func (ge *Engine) AddTopic(name string, handler func(string, interface{}), numConsumers ...int) {
	name = strings.ToLower(name)
	if _, ok := ge.topics[name]; !ok {
		if len(numConsumers) > 0 {
			ge.topics[name] = NewTopic(ge.ctx, name, handler, numConsumers[0])
		} else {
			ge.topics[name] = NewTopic(ge.ctx, name, handler)
		}
		ge.topics[name].run()
		ge.logger.Debugf("topic '%s' created", name)
	} else {
		ge.logger.Warnf("topic '%s' already exists", name)
	}
}

func (ge *Engine) Publish(name string, obj interface{}) error {
	name = strings.ToLower(name)
	if _, ok := ge.topics[name]; ok {
		return ge.topics[name].publish(newInternalMessage(obj))
	} else {
		err := fmt.Errorf("topic '%s' does not exist", name)
		ge.logger.Error(err)
		return err
	}
}
