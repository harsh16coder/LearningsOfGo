package sync

import "sync"

type counterstruct struct {
	sync.Mutex
	counter int
}

func (c *counterstruct) Count() {
	c.Lock()
	defer c.Unlock()
	c.counter++
}

func (c *counterstruct) Value() int {
	return c.counter
}
