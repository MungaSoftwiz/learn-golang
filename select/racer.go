package main

import (
	"time"
	"net/http"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// make a function called WebsiteRacer which takes two URLs and "races" them
// It hits them with an HTTP GET and returning the URL which returned first
// If none of them return within 10 seconds then it should return an error.
