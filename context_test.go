package gokaf

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func Test_setEngineKey(t *testing.T) {
	t.Run("EngineKey", func(t *testing.T) {
		engine := "engine-0"
		ctx := context.Background()
		ctx = setEngineKey(ctx, engine)
		got := getLogFields(ctx)
		if strings.Compare(got[EngineKey].(string), engine) != 0 {
			t.Errorf("engine = '%s', want '%s'", got[EngineKey].(string), engine)
		}
	})
}

func Test_setTopicKey(t *testing.T) {
	t.Run("TopicKey", func(t *testing.T) {
		topic := "topic-0"
		ctx := context.Background()
		ctx = setTopicKey(ctx, topic)
		got := getLogFields(ctx)
		if strings.Compare(got[TopicKey].(string), topic) != 0 {
			t.Errorf("topic = '%s', want '%s'", got[TopicKey].(string), topic)
		}
	})
}

func Test_setConsumerKey(t *testing.T) {
	t.Run("ConsumerKey", func(t *testing.T) {
		consumer := 0
		ctx := context.Background()
		ctx = setConsumerKey(ctx, consumer)
		got := getLogFields(ctx)
		if strings.Compare(got[ConsumerKey].(string), fmt.Sprintf("%d", consumer)) != 0 {
			t.Errorf("consumer = '%s', want '%d'", got[ConsumerKey].(string), consumer)
		}
	})
}

func Test_setProducerKey(t *testing.T) {
	t.Run("ProducerKey", func(t *testing.T) {
		ctx := context.Background()
		ctx = setProducerKey(ctx)
		got := getLogFields(ctx)
		if strings.Compare(got[ProducerKey].(string), ProducerValue) != 0 {
			t.Errorf("consumer = '%s', want '%s'", got[ProducerKey].(string), ProducerValue)
		}
	})
}

func Test_getLogFields(t *testing.T) {
	engine := "engine-0"
	topic := "topic-0"

	t.Run("LogFields for Consumer", func(t *testing.T) {
		consumer := 0
		ctx := context.Background()
		ctx = setEngineKey(ctx, engine)
		ctx = setTopicKey(ctx, topic)
		ctx = setConsumerKey(ctx, consumer)
		got := getLogFields(ctx)
		if strings.Compare(got[EngineKey].(string), engine) != 0 {
			t.Errorf("engine = '%s', want '%s'", got[EngineKey].(string), engine)
		}
		if strings.Compare(got[TopicKey].(string), topic) != 0 {
			t.Errorf("topic = '%s', want '%s'", got[TopicKey].(string), topic)
		}
		if strings.Compare(got[ConsumerKey].(string), fmt.Sprintf("%d", consumer)) != 0 {
			t.Errorf("consumer = '%s', want '%d'", got[ConsumerKey].(string), consumer)
		}
	})

	t.Run("LogFields for Producer", func(t *testing.T) {
		ctx := context.Background()
		ctx = setEngineKey(ctx, engine)
		ctx = setTopicKey(ctx, topic)
		ctx = setProducerKey(ctx)
		got := getLogFields(ctx)
		if strings.Compare(got[EngineKey].(string), engine) != 0 {
			t.Errorf("engine = '%s', want '%s'", got[EngineKey].(string), engine)
		}
		if strings.Compare(got[TopicKey].(string), topic) != 0 {
			t.Errorf("topic = '%s', want '%s'", got[TopicKey].(string), topic)
		}
		if strings.Compare(got[ProducerKey].(string), ProducerValue) != 0 {
			t.Errorf("producer = '%s', want '%s'", got[ProducerKey].(string), ProducerValue)
		}
	})
}
