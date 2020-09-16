package main

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

func SumAll(numbersToSum ...[]int) (sums []int) {
    return
}