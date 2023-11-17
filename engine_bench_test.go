// Gokaf is a simple In-memory PubSub Engine
package gokaf

import (
	"bytes"
	"fmt"
	"testing"
)

var messageBytes = []struct {
	input int
}{
	{input: 2048},   // 2kB
	{input: 32768},  // 32kB
	{input: 131072}, // 128kB
}

var topicBufferSize = []struct {
	input int
}{
	{input: 0},
	{input: 2048},
	{input: 8192},
}

func BenchmarkEngine(b *testing.B) {
	for _, ms := range messageBytes {
		mbtitle := fmt.Sprintf("msg_bytes_%dk", ms.input/1024)
		for _, bts := range topicBufferSize {
			b.Run(fmt.Sprintf("%s/buffer_topic_%dk", mbtitle, bts.input/1024), func(b *testing.B) {
				// Create a new Engine
				engine := NewEngine(mockLogger)

				topicName := "BenchTopic"
				_ = engine.RegisterTopic(topicName, bts.input)

				counter := &counter{}

				done := make(chan struct{})
				mockHandler := func(receivedMsg interface{}) {
					counter.Increment()
					if counter.Value() == b.N {
						close(done)
					}
				}

				consumer, _ := engine.GetConsumer(topicName, mockHandler)

				producer, _ := engine.GetProducer(topicName)

				consumer.Run()

				// Run the benchmark
				b.ResetTimer()

				go func() {
					for i := 0; i < b.N; i++ {
						go func() {
							// Publish messages to topics
							_ = producer.Publish(bytes.Repeat([]byte{'1'}, ms.input))
						}()
					}
				}()

				<-done
				engine.Stop()
			})
		}
	}
}
