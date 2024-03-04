package main

import (
	"sync"
	"testing"
)

// API to give us a method to increment the counter & retrieve it's value
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// sync package provides basic synchronization primitives
		// Some are: sync.Mutex, sync.WaitGroup, sync.atomic
		// these primitives coordinate execution of goroutines
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		/* test fails because multiple goroutines are trying to
		 * mutate the value of the counter at the same time
		 * so we'll change our go code to make it work
		 */
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}


// we change our signature cause "a mutex must not be copied after first use"
func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
