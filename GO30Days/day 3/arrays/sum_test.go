package main

import "testing"

// test case for function Sum

func TestSum(t *testing.T) {

    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}

        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

}

// Sum: take numbers in arrays and sum it altogether
// range loops though an arrays by return the value and index
// blank identifier used to ignore index as in for range loops

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers { 
		sum += number 				
	}
	return sum
}
