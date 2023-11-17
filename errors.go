package gokaf

import (
	"fmt"
)

// TopicClosedError is a custom error type for indicating that a topic is already closed.
type TopicClosedError struct {
	topicName string
}

// Error returns a string representation of the TopicClosedError.
func (e TopicClosedError) Error() string {
	return fmt.Sprintf("Topic %s is already closed", e.topicName)
}

// newTopicClosedError creates a new instance of TopicClosedError.
func newTopicClosedError(topicName string) error {
	return TopicClosedError{topicName: topicName}
}

// TopicExistsError is a custom error type for indicating that a topic already exists.
type TopicExistsError struct {
	topicName string
}

// Error returns a string representation of the TopicExistsError.
func (e TopicExistsError) Error() string {
	return fmt.Sprintf("Topic %s already exists", e.topicName)
}

// newTopicExistsError creates a new instance of TopicExistsError.
func newTopicExistsError(topicName string) error {
	return TopicExistsError{topicName: topicName}
}
