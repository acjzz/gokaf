/*
gokaf is a robust in-memory pubsub engine meticulously crafted to provide seamless and nearly real-time data streams.
*/
package gokaf

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Engine struct represents the core of a pubsub system, managing topics and their associated consumers.
type Engine struct {
	ctx      context.Context
	cancel   context.CancelFunc
	sigChan  chan os.Signal
	logger   *slog.Logger
	wg       sync.WaitGroup // Add a WaitGroup for synchronization
	topics   map[string]*topic
	muTopics sync.RWMutex
}

// NewEngine creates a new Engine instance, initializing the necessary components for the pubsub system.
func NewEngine(logger *slog.Logger) *Engine {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)

	e := Engine{
		ctx,
		cancel,
		sigChan,
		logger,
		sync.WaitGroup{},
		make(map[string]*topic),
		sync.RWMutex{},
	}

	e.wg.Add(1) // Increment the WaitGroup counter

	// Notify the signalChan for specified signals (interrupt and termination)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		defer close(sigChan)
		// Wait for a signal
		sig := <-sigChan

		e.logger.Error(fmt.Sprintf("Received signal: %v", sig))

		// Perform cleanup or shutdown operations here
		e.Stop()
	}()

	return &e
}

// Stop terminates the pubsub engine, initiating a graceful shutdown process.
func (e *Engine) Stop() {
	defer e.wg.Done() // Decrement the WaitGroup counter when the goroutine completes

	// Shutdown. Cancel application context will kill all attached tasks.
	e.logger.Warn("Engine Shutting Down")
	e.cancel()
}

// TopicExists checks if a specified topic exists within the pubsub engine.
func (e *Engine) TopicExists(topicName string) bool {
	e.muTopics.RLock()
	defer e.muTopics.RUnlock()

	_, exists := e.topics[topicName]
	return exists
}

// RegisterTopic creates and registers a new topic within the pubsub engine with the specified name and buffer size.
func (e *Engine) RegisterTopic(topicName string, bufferSize int) error {
	if e.TopicExists(topicName) {
		return newTopicExistsError(topicName)
	}
	e.muTopics.Lock()
	defer e.muTopics.Unlock()
	// Register the topic
	e.topics[topicName] = newTopic(e.ctx, e.logger, topicName, bufferSize)
	return nil
}

// GetProducer retrieves or creates a new producer associated with the specified topic in the pubsub engine.
func (e *Engine) GetProducer(topicName string) (*Producer, error) {
	if !e.TopicExists(topicName) {
		return nil, newTopicExistsError(topicName)
	}
	topic := e.topics[topicName]
	return newProducer(topic, e.logger), nil
}

// GetConsumer retrieves or creates a new consumer associated with the specified topic in the pubsub engine.
func (e *Engine) GetConsumer(topicName string, handler func(interface{})) (*Consumer, error) {
	if !e.TopicExists(topicName) {
		return nil, newTopicExistsError(topicName)
	}
	topic := e.topics[topicName]
	return newConsumer(topic, e.logger, handler), nil
}
