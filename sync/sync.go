package main

// we keep some state for the counter in our datatype and increment on every Inc call
type Counter struct {
	value int
}

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
