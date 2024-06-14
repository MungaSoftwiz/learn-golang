package reflection

import (
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	// creates an anonymous struct with single field initialized with {expected}
	//
	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		// takes string as input and appends it to got slice
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}
