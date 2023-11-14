// Gokaf is a simple In-memory PubSub Engine
package gokaf

import (
	"fmt"
	"testing"
	"time"
)

// MockLogger is a simple mock implementation of the Logger interface for testing.
type mockLogger struct {
	Logs []string
}

func (m *mockLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
	message := fmt.Sprintf(format, v...)
	m.Logs = append(m.Logs, message)
}

func TestEngine(t *testing.T) {
	// Create a mock logger for testing
	ml := &mockLogger{}

	// Create a new Engine with the mock logger
	e := NewEngine(ml)

	// Create channels for subscribers
	subscriber1 := make(chan interface{})
	subscriber2 := make(chan interface{})

	// Subscribe channels to topics
	e.Subscribe("news", subscriber1)
	e.Subscribe("sports", subscriber2)

	// Register handlers for specific topics
	e.AddHandler("news", func(message interface{}) {})
	e.AddHandler("sports", func(message interface{}) {})

	// Publish messages to topics
	e.Publish("news", "Breaking news: Go is awesome!")
	e.Publish("sports", map[string]int{"score": 42, "player": 7})

	// Allow some time for messages to be processed
	time.Sleep(100 * time.Millisecond)

	// Check if logs contain expected entries
	expectedLogs := []string{
		"Subscribed channel to topic: news",
		"Subscribed channel to topic: sports",
		"Added handler for topic: news",
		"Added handler for topic: sports",
		"Published message to topic: news",
		"Published message to topic: sports",
	}
	for i, expectedLog := range expectedLogs {
		if i >= len(ml.Logs) {
			t.Errorf("Expected log entry missing: %s", expectedLog)
			break
		}
		if ml.Logs[i] != expectedLog {
			t.Errorf("Expected log entry mismatch: got %s, expected %s", ml.Logs[i], expectedLog)
		}
	}

	// Unsubscribe channels from topics
	e.Unsubscribe("news", subscriber1)
	e.Unsubscribe("sports", subscriber2)

	// Allow some time for unsubscriptions to be processed
	time.Sleep(100 * time.Millisecond)

	// Check if logs contain expected unsubscription entries
	expectedUnsubscribeLogs := []string{
		"Unsubscribed channel from topic: news",
		"Unsubscribed channel from topic: sports",
	}
	for i, expectedLog := range expectedUnsubscribeLogs {
		index := len(ml.Logs) - len(expectedUnsubscribeLogs) + i
		if index < 0 {
			t.Errorf("Expected unsubscription log entry missing: %s", expectedLog)
			break
		}
		if ml.Logs[index] != expectedLog {
			t.Errorf("Expected unsubscription log entry mismatch: got %s, expected %s", ml.Logs[index], expectedLog)
		}
	}
}
