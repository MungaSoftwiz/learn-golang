package main

import (
	"time"
	"net/http"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}

// make a function called WebsiteRacer which takes two URLs and "races" them
// It hits them with an HTTP GET and returning the URL which returned first
// If none of them return within 10 seconds then it should return an error.
