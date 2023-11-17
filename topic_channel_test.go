package gokaf

import (
	"testing"
)

func TestTopicChannel(t *testing.T) {
	// Test newTopicChannel function
	bufferSize := 10
	sc := newTopicChannel(bufferSize)

	// Test if the channel is not closed initially
	if sc.IsClosed() {
		t.Error("Expected topic channel to be open initially, but it is closed.")
	}

	// Test Close method
	sc.Close()

	// Test if the channel is closed after calling Close
	if !sc.IsClosed() {
		t.Error("Expected topic channel to be closed after calling Close, but it is still open.")
	}

	// Test that closing the channel again has no effect
	sc.Close()

	// Test sending to a closed channel panics
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected sending to a closed channel to panic, but it did not.")
		}
	}()
	sc.ch <- "test"
}

func TestTopicChannelWithBuffering(t *testing.T) {
	// Test newTopicChannel function with buffering
	bufferSize := 2
	sc := newTopicChannel(bufferSize)

	// Test sending messages to the channel
	sc.ch <- "message1"
	sc.ch <- "message2"

	// Test receiving messages from the channel
	msg1 := <-sc.ch
	msg2 := <-sc.ch

	// Test if the received messages match the sent messages
	if msg1 != "message1" || msg2 != "message2" {
		t.Error("Received messages do not match the sent messages.")
	}

	// Test closing the channel with buffered messages
	sc.Close()

	// Test if the channel is closed after calling Close
	if !sc.IsClosed() {
		t.Error("Expected topic channel to be closed after calling Close, but it is still open.")
	}
}
