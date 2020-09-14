package main

import "fmt"

// Credit to sentdex

func main() {
	x := 15
	a := &x // memory address

	fmt.Println("Memory Address is", a)
	fmt.Println("Value is", *a)	// use * to read the value at that memory address

	*a = 5	// *a (or x) was added new value
	fmt.Println("Value is", x)
	fmt.Println("Memory Address is", a)

	*a = *a**a  // this equipvalent of x multiply x or (*a)(*a) => *a
	fmt.Println("return of x is", x)
	fmt.Println("return of *a is", *a)

}