package main

import (
	gofka "./src"
	"fmt"
	"time"
)

func main(){
	ge := gofka.NewGofkaEngine("test")
	topicName := "test"
	ge.AddTopic(topicName)
	i := 0
	for {
		e := ge.Publish(topicName, gofka.NewInternalMessage(fmt.Sprintf("Message%d", i)))
		if e != nil {
			fmt.Println(e)
			break
		}
		i += 1
		time.Sleep(time.Second)
	}
}
