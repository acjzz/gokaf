package gokaf

import (
	"context"
	"testing"
	"time"
)

func TestConsumerClose(t *testing.T) {
	topicName := "testTopic"
	topic := newTopic(context.Background(), mockLogger, topicName, 0)

	// Mock handler function for the consumer
	mockHandler := func(interface{}) {}

	consumer := newConsumer(topic, mockLogger, mockHandler)

	// Test: Consumer close
	t.Run("ConsumerClose", func(t *testing.T) {
		consumer.Close()

		// Wait for the Consumer to finish closing (WaitGroup counter to reach zero)
		done := make(chan struct{})
		go func() {
			defer close(done)
			consumer.wg.Wait()
		}()

		select {
		case <-done:
			// The consumer has finished closing
		case <-time.After(time.Second):
			t.Error("Timed out waiting for consumer to close")
		}
	})
}

func TestConsumerRun(t *testing.T) {
	topicName := "testTopic"
	topic := newTopic(context.Background(), mockLogger, topicName, 0)

	sentMsg := "Go is awesome"
	done := make(chan struct{})

	// Mock handler function for the consumer
	mockHandler := func(receivedMsg interface{}) {
		defer close(done)
		if receivedMsg != sentMsg {
			t.Errorf("Expected message %v, got %v", sentMsg, receivedMsg)
		}
	}

	consumer := newConsumer(topic, mockLogger, mockHandler)

	// Test: Consumer close
	t.Run("ConsumerRun", func(t *testing.T) {
		consumer.Run()

		go func() {
			consumer.topic.channel.ch <- sentMsg
		}()

		select {
		case <-done:
			// The consumer has finished closing
		case <-time.After(time.Second):
			t.Error("Timed out waiting for consumer to close")
		}
	})
}

func TestConsumerStopAfterRun(t *testing.T) {
	topicName := "testTopic"
	topic := newTopic(context.Background(), mockLogger, topicName, 0)

	sentMsg := "Go is awesome"
	done := make(chan struct{})

	// Mock handler function for the consumer
	mockHandler := func(receivedMsg interface{}) {
		close(done)
		if receivedMsg != sentMsg {
			t.Errorf("Expected message %v, got %v", sentMsg, receivedMsg)
		}
	}

	consumer := newConsumer(topic, mockLogger, mockHandler)

	// Test: Consumer close
	t.Run("ConsumerRun", func(t *testing.T) {
		consumer.Run()

		go func() {
			consumer.topic.channel.ch <- sentMsg
		}()

		select {
		case <-done:
			consumer.Close()
			consumer.wg.Wait()
		case <-time.After(time.Second):
			t.Error("Timed out waiting for consumer to close")
		}
	})
}
