package gokaf

import (
	"context"
	"testing"
	"time"
)

func TestProducerClose(t *testing.T) {
	topicName := "testTopic"
	topic := newTopic(context.Background(), mockLogger, topicName, 0)

	producer := newProducer(topic, mockLogger)

	// Test: Producer close
	t.Run("ProducerClose", func(t *testing.T) {
		producer.Stop()

		// Wait for the producer to finish closing (WaitGroup counter to reach zero)
		done := make(chan struct{})
		go func() {
			defer close(done)
			producer.wg.Wait()
		}()

		select {
		case <-done:
			// The producer has finished closing
		case <-time.After(time.Second):
			t.Error("Timed out waiting for producer to close")
		}
	})
}

func TestProducerPublish(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	topicName := "testTopic"
	topic := newTopic(ctx, mockLogger, topicName, 0)
	producer := newProducer(topic, mockLogger)

	// Use a channel to capture published messages
	publishedMessages := make(chan interface{}, 3)

	// Start a goroutine to simulate message consumption
	go func() {
		for msg := range producer.topic.channel.ch {
			publishedMessages <- msg
		}
	}()

	message := "testMessage"

	// Publish message
	err := producer.Publish(message)
	if err != nil {
		t.Errorf("Error publishing message1: %v", err)
	}

	producer.Stop()
	producer.wg.Wait()

	topic.close()
	topic.wg.Wait()

	select {
	case receivedMsg := <-publishedMessages:
		if receivedMsg != message {
			t.Errorf("Expected message %v, got %v", message, receivedMsg)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for message")
	}
}
