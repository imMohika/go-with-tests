package counter

import "sync"

type Counter struct {
	mu      sync.Mutex
	current int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.current += 1
}

func (c *Counter) Value() int {
	return c.current
}
