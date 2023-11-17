package gokaf

import (
	"sync"
)

type topicChannel struct {
	ch     chan interface{}
	closed bool
	mu     sync.Mutex
}

func newTopicChannel(bufferSize int) *topicChannel {
	return &topicChannel{
		ch:     make(chan interface{}, bufferSize),
		closed: false,
	}
}

func (sc *topicChannel) Close() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	if !sc.closed {
		close(sc.ch)
		sc.closed = true
	}
}

func (sc *topicChannel) IsClosed() bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.closed
}
