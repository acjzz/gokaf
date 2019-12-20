package main

import (
	gofka "./src"
	"fmt"
	"time"
)

func main(){
	ge := gofka.NewGofkaEngine("Engine")
	topicName := "Topic0"
	ge.AddTopic(topicName)
	i := 0
	for {
		e := ge.Publish(topicName, fmt.Sprintf("Message%d", i))
		if e != nil {
			break
		}
		i += 1
		time.Sleep(time.Second)
	}
}
