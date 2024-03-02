package main

import (
	"testing"
)

// API to give us a method to increment the counter & retrieve it's value
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d want %d", counter.Value(), 3)
		}
	})
}
