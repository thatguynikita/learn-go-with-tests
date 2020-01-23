package main

import "sync"

// Counter holds number of counts
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns new Counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter by one
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value prints the value of the counter
func (c *Counter) Value() int {
	return c.value
}
