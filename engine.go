// Gokaf is a simple In-memory PubSub Engine
package gokaf

import (
	"sync"
)

// Handler is a function type for handling messages.
type Handler func(interface{})

// Gokaf PubSub Engine
type Engine struct {
	subscribers map[string]map[chan interface{}]struct{}
	handlers    map[string][]Handler
	logger      Logger
	mu          sync.RWMutex
}

// NewEngine creates a new instance of the Engine.
func NewEngine(logger Logger) *Engine {
	return &Engine{
		subscribers: make(map[string]map[chan interface{}]struct{}),
		handlers:    make(map[string][]Handler),
		logger:      logger,
		mu:          sync.RWMutex{},
	}
}

// Subscribes a channel to a topic.
func (e *Engine) Subscribe(topic string, ch chan interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.subscribers[topic]; !exists {
		e.subscribers[topic] = make(map[chan interface{}]struct{})
	}

	e.subscribers[topic][ch] = struct{}{}
	e.logger.Printf("Subscribed channel to topic: %s", topic)
}

// Unsubscribes a channel from a topic.
func (e *Engine) Unsubscribe(topic string, ch chan interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if subscribers, exists := e.subscribers[topic]; exists {
		delete(subscribers, ch)
		if len(subscribers) == 0 {
			delete(e.subscribers, topic)
		}
		e.logger.Printf("Unsubscribed channel from topic: %s", topic)
	}
}

// AddHandler adds a handler function for a specific topic.
func (e *Engine) AddHandler(topic string, handler Handler) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.handlers[topic] = append(e.handlers[topic], handler)
	e.logger.Printf("Added handler for topic: %s", topic)
}

// Publishes a message to a topic, broadcasting it to all subscribers
// and calling the registered handlers for the topic.
func (e *Engine) Publish(topic string, message interface{}) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if subscribers, exists := e.subscribers[topic]; exists {
		for ch := range subscribers {
			go func(c chan interface{}) {
				c <- message
			}(ch)
		}
		e.logger.Printf("Published message to topic: %s", topic)
	}

	if handlers, exists := e.handlers[topic]; exists {
		for _, handler := range handlers {
			go handler(message)
		}
	}
}
