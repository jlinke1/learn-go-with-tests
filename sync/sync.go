package sync

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}
