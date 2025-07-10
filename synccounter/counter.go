package sync

import (
	"sync"
)

type counterstruct struct {
	mu      sync.Mutex
	Counter int
}

func NewCounter() *counterstruct {
	return &counterstruct{}
}
func (c *counterstruct) Count() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Counter++
}

func (c *counterstruct) Value() int {
	return c.Counter
}
