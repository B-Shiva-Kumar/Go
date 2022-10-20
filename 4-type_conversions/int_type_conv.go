package main

import "fmt"

func main() {

	// NOTE: we have to do type conversion ecxplicitly

	var (
		num1 int8  = 10
		num2 int16 = 20

		// sum12 = num1 + num2
		// err: mismatched types int8 and int

		// sum12 = num1 + int16(num2)
		// err: mismatched types int8 and int16

		sum12 = num1 + int8(num2)
	)

	fmt.Printf("Sum12 value: %v", sum12)
}
