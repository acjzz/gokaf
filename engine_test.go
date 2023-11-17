package gokaf

import (
	"os"
	"testing"
	"time"
)

func TestEngineStop(t *testing.T) {
	// Create a new Engine with a mock logger
	engine := NewEngine(mockLogger)

	// Simulate an interrupt signal to trigger the Stop method
	go func() {
		time.Sleep(100 * time.Millisecond)
		engine.Stop()
	}()

	engine.wg.Wait()

	// Check if the logger was called with the expected message
	expectedMsg := "Engine Shutting Down"
	logs := getMockLogs()
	last := logs[len(logs)-1]

	if last != expectedMsg {
		t.Errorf("Expected log message: %s, got: %s", expectedMsg, last)
	}
}

func TestInterruptSignalHandling(t *testing.T) {
	// Create a new Engine with a mock logger
	engine := NewEngine(mockLogger)

	// Simulate sending a kill signal to the Engine after a short delay
	go func() {
		time.Sleep(100 * time.Millisecond)
		engine.sigChan <- os.Interrupt
	}()

	engine.wg.Wait()

	// Check if the logger was called with the expected message
	logs := getMockLogs()

	expectedMsg := "Received signal: interrupt"
	secondLast := logs[len(logs)-2]

	if secondLast != expectedMsg {
		t.Errorf("Expected log message: %s, got: %s", expectedMsg, secondLast)
	}

	expectedMsg = "Engine Shutting Down"
	last := logs[len(logs)-1]

	if last != expectedMsg {
		t.Errorf("Expected log message: %s, got: %s", expectedMsg, last)
	}
}

func TestTopicExists(t *testing.T) {
	engine := NewEngine(mockLogger)

	// Example: Register a topic
	topicName := "exampleTopic"
	err := engine.RegisterTopic(topicName, 0)

	if err != nil {
		t.Errorf("Unexpected error while registering a new topic: %v", err)
	}
	// Test: Check if a topic exists
	t.Run("ExistingTopic", func(t *testing.T) {
		if !engine.TopicExists(topicName) {
			t.Error("Expected the topic to exist, but it doesn't.")
		}
	})

	// Test: Check if nonExistingTopic exists
	t.Run("NonExistingTopic", func(t *testing.T) {
		if engine.TopicExists("nonExistingTopic") {
			t.Error("Expected the topic not to exist, but it does.")
		}
	})
}

func TestRegisterTopic(t *testing.T) {
	engine := NewEngine(mockLogger)

	// Test: Register a new topic
	t.Run("RegisterNewTopic", func(t *testing.T) {
		topicName := "newTopic"
		err := engine.RegisterTopic(topicName, 0)

		if err != nil {
			t.Errorf("Unexpected error while registering a new topic: %v", err)
		}

		if !engine.TopicExists(topicName) {
			t.Error("Expected the newly registered topic to exist, but it doesn't.")
		}
	})

	// Test: Attempt to register an existing topic
	t.Run("RegisterExistingTopic", func(t *testing.T) {
		topicName := "existingTopic"
		err := engine.RegisterTopic(topicName, 0)

		if err != nil {
			t.Errorf("Unexpected error while registering a new topic: %v", err)
		}

		err = engine.RegisterTopic(topicName, 0)

		if err == nil {
			t.Error("Expected an error when attempting to register an existing topic, but got nil.")
		}

		expectedError := newTopicExistsError(topicName).Error()
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})
}

func TestGetProducer(t *testing.T) {
	// Create a new instance of your Engine
	engine := NewEngine(mockLogger)

	t.Run("GetProducerNoneExistingTopic", func(t *testing.T) {
		// Test case 1: Non-existing topic
		topicName := "nonexistentTopic"
		producer1, err1 := engine.GetProducer(topicName)

		if err1 == nil {
			t.Errorf("Unexpected error for non-existing topic: %v", err1)
		}

		if producer1 != nil {
			t.Error("Expected a nil producer for non-existing topic, but got non-nil")
		}
	})

	t.Run("GetProducerExistingTopic", func(t *testing.T) {
		// Test case 2: Existing topic
		existingTopic := "existingTopic"

		err := engine.RegisterTopic(existingTopic, 0)

		if err != nil {
			t.Errorf("Unexpected error while registering a new topic: %v", err)
		}

		producer2, err2 := engine.GetProducer(existingTopic)

		if err2 != nil {
			t.Errorf("Expected error for existing topic %s: %v", existingTopic, err2)
		}

		if producer2 == nil {
			t.Error("Expected a producer for existing topic, but got nil")
		}
	})
}

func TestGetConsumer(t *testing.T) {
	// Create a new instance of your Engine
	engine := NewEngine(mockLogger)

	// Mock handler function for the consumer
	mockHandler := func(interface{}) {}

	t.Run("GetConsumerNoneExistingTopic", func(t *testing.T) {
		// Test case 1: Non-existing topic
		topicName := "nonexistentTopic"

		consumer1, err1 := engine.GetConsumer(topicName, mockHandler)

		if err1 == nil {
			t.Errorf("Unexpected error for non-existing topic: %v", err1)
		}

		if consumer1 != nil {
			t.Error("Expected a nil consumer for non-existing topic, but got non-nil")
		}
	})

	t.Run("GetConsumerExistingTopic", func(t *testing.T) {
		// Test case 2: Existing topic
		existingTopic := "existingTopic"

		err := engine.RegisterTopic(existingTopic, 0)

		if err != nil {
			t.Errorf("Unexpected error while registering a new topic: %v", err)
		}

		consumer2, err2 := engine.GetConsumer(existingTopic, mockHandler)

		if err2 != nil {
			t.Errorf("Expected error for existing topic %s: %v", existingTopic, err2)
		}

		if consumer2 == nil {
			t.Error("Expected a consumer for existing topic, but got nil")
		}
	})
}

func TestProducerConsumer(t *testing.T) {
	// Create a new instance of your Engine
	engine := NewEngine(mockLogger)

	topicName := "testTopic"
	err := engine.RegisterTopic(topicName, 0)

	if err != nil {
		t.Errorf("Unexpected error while registering a new topic: %v", err)
	}

	sentMsg := "Go is awesome"
	done := make(chan struct{})

	// Mock handler function for the consumer
	mockHandler := func(receivedMsg interface{}) {
		defer close(done)
		if receivedMsg != sentMsg {
			t.Errorf("Expected message %v, got %v", sentMsg, receivedMsg)
		}
	}

	consumer, err1 := engine.GetConsumer(topicName, mockHandler)

	if err1 != nil {
		t.Errorf("Unexpected error while getting a consumer for topic %s: %v", topicName, err1)
	}

	producer, err2 := engine.GetProducer(topicName)

	if err2 != nil {
		t.Errorf("Unexpected error while getting a producer for topic %s: %v", topicName, err2)
	}

	consumer.Run()

	go func() {
		err3 := producer.Publish(sentMsg)

		if err3 != nil {
			t.Errorf("Error publishing message1: %v", err3)
		}
	}()

	select {
	case <-done:
		// The consumer has finished closing
	case <-time.After(time.Second):
		t.Error("Timed out waiting for consumer to close")
	}
}
