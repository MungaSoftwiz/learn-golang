package concurrency

import (
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)


	// We normally wait for a func to return/process to finish(It blocking)
	// An op that doesn't block will run in a separate process called a "goroutine"
	// Only way to start a goroutine is to put go im front of a function call
	// Below uses an anonymous function

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	// We give each anony func a param cause it was taking a reference to the "url" var.
	// u is a copy of the value of url & fixed as value of url for iteration
	//this "()"makes goroutine execute the same time they are declared


	time.Sleep(2 * time.Second)
	return results
}
