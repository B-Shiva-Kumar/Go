package main

import "fmt"

func main() {

	// NOTE:
	// - explicitly we have to convert a value wrt type
	// - using const -> by default it is implicitly converts the value

	const num1 = 10
	var num2 int8 = 10

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Printf("%v %T\n", num2, num2)
	fmt.Printf("%v %T\n", num1+num2, num1+num2)

}
