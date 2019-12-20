package src

import (
	"context"
	"fmt"
	"strings"
)

const (
	ENGINE_KEY = "engine-name"
	TOPIC_KEY = "topic-name"
	CONSUMER_KEY = "consumer-id"
	PRODUCER_KEY = "producer-id"
)

func setStrContextKey(ctx context.Context, key string, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func setEngineKey(ctx context.Context, value string) context.Context {
	return setStrContextKey(ctx, ENGINE_KEY, fmt.Sprintf("GE[%s]",value))
}

func setTopicKey(ctx context.Context, value string) context.Context {
	return setStrContextKey(ctx, TOPIC_KEY, fmt.Sprintf("T[%s]",value))
}

func setConsumerKey(ctx context.Context, value int) context.Context {
	return setStrContextKey(ctx, CONSUMER_KEY, fmt.Sprintf("Consumer%d",value))
}

func setProducerKey(ctx context.Context) context.Context {
	return setStrContextKey(ctx, PRODUCER_KEY, "Producer")
}

func getEngineToken(ctx context.Context) string {
	var elements []string
	for _, k := range []string{ ENGINE_KEY, TOPIC_KEY , CONSUMER_KEY, PRODUCER_KEY} {
		v := ctx.Value(k)
		if v != nil {
			elements = append(elements, v.(string))
		}
	}
	return strings.Join(elements[:], " - ")
}
