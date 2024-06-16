package main

import (
	"sync"
)

// we keep some state for the counter in our datatype and increment on every Inc call
// Mutex is a mutual extension lock
type Counter struct {
	mu    sync.Mutex
	value int
}

// Any goroutine calling Inc will acquire lock on Counter if they're first
// All other goroutines have to wait for it to Unlock before getting access
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

// Create a constructor which shows readers of your API that
// it would be better not to initialise the type yourself
func NewCounter() *Counter {
	return &Counter{}
}

/* sync package docs: https://pkg.go.dev/sync
* Mutex or Channels: https://go.dev/wiki/MutexOrChannel */
