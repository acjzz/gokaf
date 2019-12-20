package main

import (
	gofka "./src"
	"fmt"
	"time"
)

func handler(obj interface{}){
	fmt.Println(obj.(string))
}

func main(){
	ge := gofka.NewGofkaEngine("Engine")
	topicName := "Topic0"
	ge.AddTopic(topicName, handler, 5)
	i := 0
	for {
		e := ge.Publish(topicName, fmt.Sprintf("Message%d", i))
		if e != nil {
			break
		}
		i += 1
		time.Sleep(time.Millisecond)
	}
}
