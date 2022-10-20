package main

import "fmt"

func main() {

	var num1 int32 = 112
	fmt.Printf("Type: %T, value: %v \n", num1, num1)

	var num2 int = 112
	fmt.Printf("Type: %T, value: %v \n", num2, num2)

	// var x int8 = 500
	// cannot use 500 (untyped int constant) as int8 value in variable declaration (overflows)

	var x int8 = 50
	var y int16 = -4500
	fmt.Printf("Type: %T, value: %v \n", x, x)
	fmt.Printf("Type: %T, value: %v \n", y, y)
}
