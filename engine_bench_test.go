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

var topicsConsumersProdcuers = []struct {
	numTopics    int
	numProducers int
	numConsumers int
}{
	{numTopics: 2, numProducers: 1, numConsumers: 2},
	{numTopics: 2, numProducers: 2, numConsumers: 1},
	{numTopics: 5, numProducers: 3, numConsumers: 6},
	{numTopics: 5, numProducers: 6, numConsumers: 3},
}

func BenchmarkSingleTopic(b *testing.B) {
	for _, ms := range messageBytes {
		mbTitle := fmt.Sprintf("msg_%dkB", ms.input/1024)
		for _, bts := range topicBufferSize {
			b.Run(fmt.Sprintf("%s/buffer_%dkB", mbTitle, bts.input/1024), func(b *testing.B) {
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

func BenchmarkMultipleTopics(b *testing.B) {
	for _, ms := range messageBytes {
		mbTitle := fmt.Sprintf("msg_%dkB", ms.input/1024)
		for _, bts := range topicBufferSize {
			btsTitle := fmt.Sprintf("buffer_%dkB", bts.input/1024)
			for _, tpc := range topicsConsumersProdcuers {
				tpcTitle := fmt.Sprintf("%dtop_%dprod_%dcon", tpc.numTopics, tpc.numProducers, tpc.numConsumers)
				b.Run(fmt.Sprintf("%s/%s/%s", mbTitle, btsTitle, tpcTitle), func(b *testing.B) {
					// Create a new Engine
					engine := NewEngine(mockLogger)

					counters := []*counter{}
					dones := []chan struct{}{}

					for topicIndex := 0; topicIndex < tpc.numTopics; topicIndex++ {
						// Create a shared counter for the topic and consumers
						topicName := fmt.Sprintf("Topic%d", topicIndex)
						_ = engine.RegisterTopic(topicName, bts.input)

						counters = append(counters, &counter{})
						dones = append(dones, make(chan struct{}))

						// Initialize consumers
						for i := 0; i < tpc.numConsumers; i++ {
							go func(topicIndex int) {
								mockHandler := func(receivedMsg interface{}) {
									counters[topicIndex].Increment()
									if counters[topicIndex].Value() == tpc.numProducers*b.N {
										select {
										case <-dones[topicIndex]:
											// Channel is already closed
										default:
											// Channel is open, so close it
											close(dones[topicIndex])
										}
									}
								}
								consumer, _ := engine.GetConsumer(topicName, mockHandler)
								consumer.Run()

								<-dones[topicIndex]
							}(topicIndex)
						}
					}

					b.ResetTimer()
					for topicIndex := 0; topicIndex < tpc.numTopics; topicIndex++ {
						topicName := fmt.Sprintf("Topic%d", topicIndex)
						// Initialize producers
						for i := 0; i < tpc.numProducers; i++ {
							go func(topicName string) {
								for i := 0; i < b.N; i++ {
									// Publish messages to topics
									producer, _ := engine.GetProducer(topicName)
									_ = producer.Publish(bytes.Repeat([]byte{'1'}, ms.input))
								}
							}(topicName)
						}
					}

					for topicIndex := 0; topicIndex < tpc.numTopics; topicIndex++ {
						<-dones[topicIndex]
					}

					// Stop the engine
					engine.Stop()
				})
			}
		}
	}
}
