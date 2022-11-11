package main

import "fmt"

func main() {
	name := "Bill"

	// prints memory address of the variable
	fmt.Println(&name)
	fmt.Println(*&name)
}

// What is the &  operator used for?
// turning a value into a pointer

// When you see a * operator in front of a pointer, what will it turn the pointer into?
//  a value
