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

// Mock of our dependency for our tests to use
type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}

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
