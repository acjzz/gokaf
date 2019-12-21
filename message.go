package gokaf

type internalMessage struct {
	value	interface{}
}

func newInternalMessage(value interface{}) internalMessage {
	return internalMessage{value }
}
