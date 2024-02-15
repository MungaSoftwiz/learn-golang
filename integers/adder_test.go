package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Go examples are executed just like tests so they reflect
// what's in the code does unlike README.md examples which 
// become outated
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

}
