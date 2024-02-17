package main

import (
//	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Our dependency as an interface
// We see that from this dependency we create a mock & a real sleeper
type Sleeper interface {
	Sleep()
}

/* Mock of our dependency for our tests to use
type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}*/

// We have 2 different dependencies & we want to record all
// their operations into one list

const write = "write"
const sleep = "sleep"

// Implements both io.Writer and Sleeper, recording every call into one slice
// New mock of our dependency for our tests to use
type SpyCountdownOperations struct {
	Calls []string
}

// s is a receiver parameter of a method
// It's always the instance of the struct on which method is called
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}


<F2>
// A real sleeper which implements the interface we need
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
        time.Sleep(1 * time.Second)
}


// Fprint takes in an io.Writer(e.g *bytes.Buffer) and sends string to it
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		//time.Sleep(1 * time.Second)

		//Calling in the injected dependency
		//causes test to run fast with our mocks
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)

}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
