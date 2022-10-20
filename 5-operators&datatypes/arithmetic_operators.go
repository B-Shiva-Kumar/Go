package main

import "fmt"

func main() {

	i := 999
	j := 99

	fmt.Printf("addition: %v \n", i+j)
	fmt.Printf("subraction: %v \n", i-j)
	fmt.Printf("multiplication: %v \n", i*j)
	fmt.Printf("division: %v \n", i/j)
	fmt.Printf("modulus: %v \n", i%j)

	var (
		sum1 = 100 + 50
		sum2 = sum1 + 100
		sum3 = sum2 - sum1
	)

	fmt.Println("total sum value:", sum3)

}

// output:

// addition: 1098
// subraction: 900
// multiplication: 98901
// division: 10
// modulus: 9
// total sum value: 100
