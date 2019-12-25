# Gokaf
Gokaf is a simple In-memory PubSub Engine to enable near realtime data streams

## Example

```go
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
		ge.AddTopic(topicName, func(topic string, obj interface{}) {
			// Printf usage on the handlers is not recommended at all
			// if you pretend to achieve near realtime streams
			// In this example is for demonstration purposes only
			fmt.Printf("Consumed '%v' from topic '%s'\n", obj, topic)
		})
	}

	go func(){
		for i := 1; i <= 1000; i++ {
			// Simulation of High Frequency Data Stream
			e := ge.Publish(topics[0], fmt.Sprintf("High Frequency Message%d", i))
			if e != nil {
				fmt.Printf("publishing to topic %s, err: %v", topics[0], e)
				break
			}
			time.Sleep(time.Millisecond/100)
		}
	}()

	// Simulation of Low Frequency Data Stream
	for i := 1; i <= 35; i++ {
		e := ge.Publish(topics[1], fmt.Sprintf("Low Frequency Message%d", i))
		if e != nil {
			fmt.Printf("publishing to topic %s, err: %v", topics[1], e)
			break
		}
		time.Sleep(time.Millisecond)
	}

	ge.Stop()
}
```
