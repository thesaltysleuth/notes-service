package counter

import (
	"sync"
	"sync/atomic"
)

type MutexCounter struct {
	mu sync.Mutex
	x int64
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.x++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int64 { return c.x }


//----------------------------------------------


type AtomicCounter struct {
	x atomic.Int64
}

func (c *AtomicCounter) Inc() { c.x.Add(1) }
func (c *AtomicCounter) Value() int64 { return c.x.Load() }

