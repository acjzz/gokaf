package gokaf

import (
	"context"
	"fmt"
	"strings"
)

const (
	EngineKey   = "engine-name"
	TopicKey    = "topic-name"
	ConsumerKey = "consumer-id"
	ProducerKey = "producer-id"
)

func setStrContextKey(ctx context.Context, key string, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func setEngineKey(ctx context.Context, value string) context.Context {
	return setStrContextKey(ctx, EngineKey, value)
}

func setTopicKey(ctx context.Context, value string) context.Context {
	return setStrContextKey(ctx, TopicKey, value)
}

func setConsumerKey(ctx context.Context, value int) context.Context {
	return setStrContextKey(ctx, ConsumerKey, fmt.Sprintf("%d", value))
}

func setProducerKey(ctx context.Context) context.Context {
	return setStrContextKey(ctx, ProducerKey, "Producer")
}

func getEngineToken(ctx context.Context) string {
	var elements []string
	for _, k := range []string{EngineKey, TopicKey, ConsumerKey, ProducerKey} {
		v := ctx.Value(k)
		if v != nil {
			elements = append(elements, v.(string))
		}
	}
	return strings.Join(elements[:], " - ")
}

func getLogFields(ctx context.Context) map[string]interface{} {
	elements := map[string]interface{}{}

	for _, k := range []string{EngineKey, TopicKey, ConsumerKey, ProducerKey} {
		v := ctx.Value(k)
		if v != nil {
			elements[k] = v
		}
	}
	return elements
}
