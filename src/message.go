package src

type internalMessage struct {
	value	interface{}
}

func newInternalMessage(value interface{}) internalMessage {
	return internalMessage{value }
}

func (m *internalMessage) consume() {
	// DoSomething
}
