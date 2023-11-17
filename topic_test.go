package gokaf

import (
	"context"
	"testing"
	"time"
)

func TestNewTopic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	name := "testTopic"

	topic := newTopic(ctx, mockLogger, name, 0)

	// Add assertions to check if the topic is created correctly
	if topic == nil {
		t.Error("Expected non-nil topic, got nil")
	}
	if topic != nil && topic.name != name {
		t.Errorf("Expected topic name %s, got %s", name, topic.name)
	}
}

func TestTopicClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	name := "testTopic"
	topic := newTopic(ctx, mockLogger, name, 0)

	// Use a channel to check if the close method is called
	closedChan := make(chan struct{})
	go func() {
		topic.close()
		closedChan <- struct{}{}
	}()

	// Wait for the close method to complete
	select {
	case <-closedChan:
		// Close method completed successfully
	case <-time.After(time.Second):
		t.Error("Timeout waiting for close method to complete")
	}

	// Add assertions to check if the close method behaves as expected
	// For example, you can check if the context is canceled, etc.
	// You may need to modify these assertions based on your specific implementation.
	if topic.ctx.Err() != context.Canceled {
		t.Error("Expected canceled context after close, got", topic.ctx.Err())
	}
}

func TestTopicPublish(t *testing.T) {
	// Create a new topic for testing
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	topicName := "testTopic"
	testTopic := newTopic(ctx, mockLogger, topicName, 0)

	// Use a channel to capture published messages
	publishedMessages := make(chan interface{}, 3)

	// Start a goroutine to simulate message consumption
	go func() {
		for msg := range testTopic.channel.ch {
			publishedMessages <- msg
		}
	}()

	// Publish messages to the topic
	message1 := "Message 1"
	err1 := testTopic.publish(message1)
	if err1 != nil {
		t.Errorf("Error publishing message1: %v", err1)
	}

	message2 := "Message 2"
	err2 := testTopic.publish(message2)
	if err2 != nil {
		t.Errorf("Error publishing message2: %v", err2)
	}

	// Close the topic to stop the goroutine
	testTopic.close()

	testTopic.wg.Wait()

	// Verify the published messages
	select {
	case receivedMsg1 := <-publishedMessages:
		if receivedMsg1 != message1 {
			t.Errorf("Expected message1 %v, got %v", message1, receivedMsg1)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for message1")
	}

	select {
	case receivedMsg2 := <-publishedMessages:
		if receivedMsg2 != message2 {
			t.Errorf("Expected message2 %v, got %v", message2, receivedMsg2)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for message2")
	}
}

func TestClosedTopicPublish(t *testing.T) {
	// Create a new topic
	ctx, cancel := context.WithCancel(context.Background())

	name := "testTopic"
	topic := newTopic(ctx, mockLogger, name, 0)

	// Test publishing a message to the topic
	message := "Hello, world!"

	// Test publishing when the topic is closed
	cancel() // Close the topic
	err := topic.publish(message)
	if err == nil {
		t.Error("Expected error when publishing to a closed topic, but got none.")
	} else if err.Error() != "Topic testTopic is already closed" {
		t.Errorf("Unexpected error message. Expected 'Topic testTopic is already closed', but got '%v'", err.Error())
	}
}

func TestClosedTopicChannelPublish(t *testing.T) {
	// Create a new topic
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	name := "testTopic"
	topic := newTopic(ctx, mockLogger, name, 0)

	// Test publishing a message to the topic
	message := "Hello, world!"

	// Test publishing when the topic channel is closed
	topic.channel.Close()

	err := topic.publish(message)
	if err == nil {
		t.Error("Expected error when publishing to a closed topic, but got none.")
	} else if err.Error() != "Topic testTopic is already closed" {
		t.Errorf("Unexpected error message. Expected 'Topic testTopic is already closed', but got '%v'", err.Error())
	}
}
