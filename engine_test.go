package gokaf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"testing"
)

func TestEngine_Publish_Error(t *testing.T) {
	topicName := "test"

	type fields struct {
		name     string
		logLevel logrus.Level
	}
	type args struct {
		name string
		obj  interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantErrMsg string
		startTopic bool
	}{
		{
			"NoTopic",
			fields{"testEngine", logrus.ErrorLevel},
			args{topicName, "message"},
			true,
			fmt.Sprintf("topic %s does not exists", topicName),
			false,
		}, {
			"TopicStopped",
			fields{"testEngine", logrus.ErrorLevel},
			args{topicName, "message"},
			true,
			"topic closed",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ge := NewEngine(
				tt.fields.name,
				tt.fields.logLevel,
			)
			if tt.startTopic {
				ge.AddTopic(topicName, func(topic string, obj interface{}) {})
				ge.Stop()
			}
			err := ge.Publish(tt.args.name, tt.args.obj)
			if !tt.wantErr && err != nil {
				t.Errorf("publish() not error, wantErr %v", tt.wantErr)
			} else if err != nil {
				ErrMsg := fmt.Sprintf("%v", err)
				if strings.Compare(ErrMsg, tt.wantErrMsg) != 0 {
					t.Errorf("publish() error = '%v', wantErr '%v'", ErrMsg, tt.wantErrMsg)
				}
			}
			ge.Stop()
		})
	}
}

func TestEngine_Publish(t *testing.T) {
	testName := "publish & Consume"
	t.Run(testName, func(t *testing.T) {
		ge := NewEngine(testName, logrus.ErrorLevel)

		topicName := "topic"
		msg := "Test Message"

		ge.AddTopic(topicName, func(topic string, obj interface{}) {
			if strings.Compare(fmt.Sprintf("%v", obj), msg) != 0 {
				t.Errorf("publish() received = '%v', expected '%v'", obj, msg)
			} else if strings.Compare(topicName, topic) != 0 {
				t.Errorf("publish() received from topic '%s', expected '%s'", topic, topicName)
			}
		})

		err := ge.Publish(topicName, msg)
		if err != nil {
			t.Errorf("publish() error, wanted not Error\nError: %v", err)
		}
		ge.Stop()
	})
}

func TestEngine_Publish_Multiple_Topics(t *testing.T) {
	baseTopicName := "topic"
	baseMsg := "Test Message"

	type Message struct {
		topic   string
		message string
	}

	tests := []struct {
		name         string
		numTopics    int
		numConsumers int
	}{
		{"Two topics", 2, 1},
		{"Three topics twenty Consumers", 3, 20},
		{"Five topics", 5, 1},
		{"Five topics two Consumers", 5, 2},
		{"Twenty topics five Consumers", 20, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ge := NewEngine(tt.name, logrus.ErrorLevel)
			for i := 0; i < tt.numTopics; i++ {
				topicName := fmt.Sprintf("%s-%d", baseTopicName, i)
				ge.AddTopic(topicName, func(topic string, obj interface{}) {
					received := obj.(Message)
					if strings.Compare(received.topic, topic) != 0 {
						t.Errorf("publish() received from topic '%s', expected '%s'", received.topic, topicName)
					}
				}, tt.numConsumers)
			}

			var wg sync.WaitGroup
			for i := 0; i < tt.numTopics-1; i++ {
				wg.Add(1)
				topicName := fmt.Sprintf("%s-%d", baseTopicName, i)
				go func(topic string) {
					msg := Message{topicName, baseMsg}
					err := ge.Publish(topicName, msg)
					if err != nil {
						t.Errorf("publish() error, wanted not Error\nError: %v", err)
					}
					wg.Done()
				}(topicName)
			}
			wg.Add(1)
			topicName := fmt.Sprintf("%s-%d", baseTopicName, tt.numTopics-1)
			msg := Message{topicName, baseMsg}
			err := ge.Publish(topicName, msg)
			if err != nil {
				t.Errorf("publish() error, wanted not Error\nError: %v", err)
			}
			wg.Done()
			wg.Wait()
			ge.Stop()
		})
	}
}
