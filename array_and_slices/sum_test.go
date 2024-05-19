package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	//can be [...]int{1, 2, 3, 4, 5}
	t.Run("collection of 5 numbers", func(t *testing.T) {

		//we make it a slice, more convenient
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		// %v useful to print default format, inputs of func
		// in error message
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	//if got != want { go don't use equality ops with slices
	//we have slices std pkg, has slices.Equal in go 1.21+
	//to do a shallow compare on slices
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// test to add tails of each slice"collec. of all items except head"
func TestSumAllTails(t *testing.T) {

	//helper func are not included in test output
	//local variable scope only used here(anonymous func)
	//it adds type-safety. If dev adds checkSums(t, got,"dave" compiler will stop
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)

	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

// https://go.dev/blog/cover  -> go test -cover
// go test -coverprofile=coverage.out or go test -cover
