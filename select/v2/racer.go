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

	// select helps us synchronize comm across multiple channels easily
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


/*
In Go, a `chan struct{}` is used when you want a channel where the value that's being passed doesn't carry any meaningful data. Here are a few reasons why you might want to use it:

1. Signaling without data: Often, you just want to notify another goroutine that something has happened (e.g., an event has occurred or a job has finished), but you don't need to send any data about what happened. Using a `chan struct{}` is a clear signal to readers of your code that the channel is being used for signaling, not for passing data.

2. Zero memory usage: A `struct{}` in Go does not occupy any memory. So, a channel of `chan struct{}` also uses no memory for the values it holds. This can be beneficial if you're creating a large number of channels or storing a large number of values in a channel.

3. Closing to broadcast: A closed channel can be received from an unlimited number of times, always yielding the zero value of the channel's type immediately. If the channel's type is `struct{}`, then the zero value is just `{}`. This characteristic can be used for broadcasting a signal to an unlimited number of goroutines.

Here's an example of how it's used:

```go
done := make(chan struct{})

go func() {
    // Do some work...
    
    // Signal that the work is done.
    close(done)
}()

// Wait for the work to finish.
<-done
```

In this code, the `done` channel is being used to signal when some goroutine has finished its work. No data needs to be passed; the sending of the signal is all that matters.
*/


/*
In Go, the `select` statement is used for synchronizing communication across multiple channels. It allows a goroutine to wait on multiple communication operations (both send and receive) and proceed with the one that's ready first, providing a way to deal with situations where you need to handle multiple channels at once.

However, it's worth noting that while `select` is used with channels, and channels are often used with goroutines (which you might think of as similar to processes), `select` itself doesn't synchronize the goroutines or processes - it's used for the synchronization of operations (sends and receives) on channels.
*/