package src

import "fmt"

type messageInterface interface {
	consume()
}

type internalMessage struct {
	value	string
}

func NewInternalMessage(value string) *internalMessage {
	return &internalMessage{value }
}

func (m *internalMessage) consume() {
	fmt.Println(m.value)
}
