package sharedvars

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

type AtomicCounter struct {
	n atomic.Int64
}

func (c *AtomicCounter) Inc() {
	c.n.Add(1)
}

func (c *AtomicCounter) Value() int64 {
	return c.n.Load()
}

type OnceValue struct {
	once  sync.Once
	value string
}

func (o *OnceValue) Init(fn func() string) string {
	o.once.Do(func() {
		o.value = fn()
	})
	return o.value
}
