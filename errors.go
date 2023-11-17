package gokaf

import (
	"fmt"
)

// TopicClosedError is a custom error type for indicating that a topic is already closed.
type TopicClosedError struct {
	TopicName string
}

func (e TopicClosedError) Error() string {
	return fmt.Sprintf("Topic %s is already closed", e.TopicName)
}

// newTopicClosedError creates a new instance of TopicClosedError.
func newTopicClosedError(topicName string) error {
	return TopicClosedError{TopicName: topicName}
}

type TopicExistsError struct {
	TopicName string
}

func (e TopicExistsError) Error() string {
	return fmt.Sprintf("Topic %s already exists", e.TopicName)
}

// newTopicExistsError creates a new instance of TopicExistsError.
func newTopicExistsError(topicName string) error {
	return TopicExistsError{TopicName: topicName}
}
