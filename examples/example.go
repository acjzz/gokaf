package main

import (
	"github.com/acjzz/gokaf"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func main(){
	ge := gokaf.NewEngine("MyEngine", logrus.DebugLevel)

	topics := []string { "Topic0", "Topic1" }

	for _, topicName := range topics {
		// Register different Handler per each Topic as well as the Topics themselves
		ge.AddTopic(topicName, func(obj interface{}) {
			fmt.Printf("Consumed %v from topic %s", topicName, obj)
		})
	}

	go func(){
		// Simulation of Low Frequency Data Stream
		for i := 0; i < 1000; i++ {
			e := ge.Publish(topics[0], fmt.Sprintf("MessageA%d", i))
			if e != nil {
				fmt.Printf("publishing to topic %s", topics[0])
				break
			}
			time.Sleep(time.Millisecond/10)
		}
	}()

	for i := 0; i < 1000000; i++ {
		// Simulation of High Frequency Data Stream
		e := ge.Publish(topics[1], fmt.Sprintf("MessageA%d", i))
		if e != nil {
			fmt.Printf("publishing to topic %s", topics[1])
			break
		}
		time.Sleep(time.Nanosecond)
	}
	ge.Stop()
}
