package main

import (
	"bytes"
	"testing"
)

//bytes implement the io.Writer interface
//our tests can capture what data is being generated
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
