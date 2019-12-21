package gokaf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
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
			fields{ "testEngine", logrus.ErrorLevel, },
			args{ topicName, "message"},
			true,
			fmt.Sprintf("topic %s does not exists", topicName),
			false,
		},{
			"TopicStopped",
			fields{ "testEngine", logrus.ErrorLevel, },
			args{ topicName, "message"},
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
				ge.AddTopic(topicName, func(topic string, obj interface{}){})
				ge.Stop()
			}
			err := ge.Publish(tt.args.name, tt.args.obj)
			if !tt.wantErr && err != nil {
				t.Errorf("Publish() not error, wantErr %v", tt.wantErr)
			} else if err != nil {
				ErrMsg := fmt.Sprintf("%v", err)
				if strings.Compare(ErrMsg, tt.wantErrMsg) != 0 {
					t.Errorf("Publish() error = '%v', wantErr '%v'", ErrMsg, tt.wantErrMsg)
				}
			}
			ge.Stop()
		})
	}
}

func TestEngine_Publish(t *testing.T) {
	t.Run("Publish & Consume", func(t *testing.T) {
		ge := NewEngine("test2", logrus.ErrorLevel, )

		topicName := "topic"
		msg := "Test Message"

		ge.AddTopic(topicName, func(topic string, obj interface{}){
			if strings.Compare(fmt.Sprintf("%v", obj), msg) != 0 {
				t.Errorf("Publish() received = '%v', expected '%v'", obj, msg)
			} else if strings.Compare(topicName, topic) != 0 {
				t.Errorf("Publish() received from topic '%s', expected '%s'", topic, topicName)
			}
		})

		err := ge.Publish(topicName, msg)
		if err != nil {
			t.Errorf("Publish() error, wanted not Error")
		}
		ge.Stop()
	})
}
