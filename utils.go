package gokaf

import "sync"

// counter represents a simple counter.
type counter struct {
	mu    sync.RWMutex
	value int
}

// Increment increments the counter by 1.
func (c *counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Decrement decrements the counter by 1.
func (c *counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

// Decrement decrements the counter by 1.
func (c *counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
