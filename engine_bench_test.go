// Gokaf is a simple In-memory PubSub Engine
package gokaf

import (
	"testing"
)

type NullLogger struct {
}

func (n *NullLogger) Printf(format string, v ...interface{}) {
}

func BenchmarkEngine(b *testing.B) {
	// Create a new Engine
	nullLogger := &NullLogger{}

	Engine := NewEngine(nullLogger)

	// Create channels for subscribers
	subscribers := make([]chan interface{}, b.N)
	for i := 0; i < b.N; i++ {
		subscribers[i] = make(chan interface{})
	}

	// Subscribe channels to topics
	for _, ch := range subscribers {
		Engine.Subscribe("news", ch)
	}

	// Register handlers for specific topics
	Engine.AddHandler("news", func(message interface{}) {})

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Publish messages to topics
		Engine.Publish("news", "Breaking news: Go is awesome!")
	}

	// Unsubscribe channels from topics
	for _, ch := range subscribers {
		Engine.Unsubscribe("news", ch)
	}
}
