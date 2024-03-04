package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

// Lock() will block the counter struct to changes until the Unlock() is called, making it impossible
// for different goroutines to change the state of counter at same time.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
