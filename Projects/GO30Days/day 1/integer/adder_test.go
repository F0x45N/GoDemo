package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but was '%d'", expected, sum)
	}
}

// Add: take two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 4)
	fmt.Println(sum)
	// Output: 5
}
