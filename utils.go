package gokaf

import "sync"

// Counter represents a simple counter.
type Counter struct {
	mu    sync.RWMutex
	value int
}

// Increment increments the counter by 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Decrement decrements the counter by 1.
func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

// Decrement decrements the counter by 1.
func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
