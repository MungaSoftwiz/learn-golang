package concurrency

import (
	"testing"
	"time"
)
// We use a benchmark to test the speed of CheckWebsites

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	// Tests using a slice of 100 urls & uses fake WebsiteChecker implementaton
	// make([]T, length, capacity), make(map[K]V, capacity)
	// Creates slice with initial len & capacity of 100
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer() //resets time before it actually runs
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
