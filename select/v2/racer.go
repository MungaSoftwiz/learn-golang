package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {

	// select helps us synchronize processes easily
	// We now don't check for website's speed one after the other
	// select allows you to wait on multiple channels. the one that send 1st wins
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// time.After() comes very handy in select
	// It returns a chan & will send signal to it after defined time
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// Ping creates a chan struct{} and returns it
// Why struct{}? it's the smallest data type available from memory perspective
// We get no allocation Vs something like bool
// We are not sending anything to chan so why allocate space?
func ping(url string) chan struct{} {
	ch := make(chan struct{}) //always use make()
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
