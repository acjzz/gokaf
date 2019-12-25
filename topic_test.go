package gokaf

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestTopic_Publish(t *testing.T) {
	topicName := "test"
	msg := internalMessage{"test"}

	t.Run("Topic_Publish", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		topic := NewTopic(
			ctx,
			topicName,
			func(topic string, obj interface{}){
				if strings.Compare(fmt.Sprintf("%v", obj), msg.value.(string)) != 0 {
					t.Errorf("publish() received = '%v', expected '%v'", obj, msg)
				} else if strings.Compare(topicName, topic) != 0 {
					t.Errorf("publish() received from topic '%s', expected '%s'", topic, topicName)
				}
			},
		)

		topic.run()
		err := topic.publish(msg)
		if err != nil {
			t.Errorf("publish() err: %v", err)
		}
		cancel()
	})
}
