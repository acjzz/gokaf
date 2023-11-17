package gokaf

import (
	"fmt"
	"testing"
)

func TestTopicClosedError(t *testing.T) {
	// Test: Error message construction
	t.Run("ErrorMessageConstruction", func(t *testing.T) {
		topicName := "closedTopic"
		err := newTopicClosedError(topicName)

		expectedErrorMessage := fmt.Sprintf("Topic %s is already closed", topicName)
		actualErrorMessage := err.Error()

		if actualErrorMessage != expectedErrorMessage {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, actualErrorMessage)
		}
	})
}

func TestTopicExistsError(t *testing.T) {
	// Test: Error message construction
	t.Run("ErrorMessageConstruction", func(t *testing.T) {
		topicName := "existingTopic"
		err := newTopicExistsError(topicName)

		expectedErrorMessage := fmt.Sprintf("Topic %s already exists", topicName)
		actualErrorMessage := err.Error()

		if actualErrorMessage != expectedErrorMessage {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, actualErrorMessage)
		}
	})
}
