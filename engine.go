package gokaf

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"sync"
)

type Engine struct {
	ctx         context.Context
	ctxCancel   context.CancelFunc
	logger      LogWrapper
	topics      map[string]*Topic
	topicsMutex sync.RWMutex
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
		sync.RWMutex{},
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

func (ge *Engine) doesTopicExists(name string) bool {
	ge.topicsMutex.RLock()
	_, ok := ge.topics[name]
	ge.topicsMutex.RUnlock()
	return ok
}

func (ge *Engine) AddTopic(name string, handler func(string, interface{}), numConsumers ...int) {
	name = strings.ToLower(name)
	if !ge.doesTopicExists(name) {
		if len(numConsumers) > 0 {
			ge.topicsMutex.Lock()
			ge.topics[name] = NewTopic(ge.ctx, name, handler, numConsumers[0])
			ge.topicsMutex.Unlock()
		} else {
			ge.topicsMutex.Lock()
			ge.topics[name] = NewTopic(ge.ctx, name, handler)
			ge.topicsMutex.Unlock()
		}
		ge.topicsMutex.RLock()
		ge.topics[name].run()
		ge.topicsMutex.RUnlock()
		ge.logger.Debugf("topic '%s' created", name)
	} else {
		ge.logger.Warnf("topic '%s' already exists", name)
	}
}

func (ge *Engine) Publish(name string, obj interface{}) error {
	name = strings.ToLower(name)
	if ge.doesTopicExists(name) {
		ge.topicsMutex.RLock()
		err := ge.topics[name].publish(newInternalMessage(obj))
		ge.topicsMutex.RUnlock()
		return err
	} else {
		err := fmt.Errorf("topic '%s' does not exist", name)
		ge.logger.Error(err)
		return err
	}
}
