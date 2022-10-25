package main

import "fmt"

func main() {

	// 1
	//  {} -> initializer
	marks := [3]int{80, 90, 100}
	fmt.Printf("value: %v, type: %T\n", marks, marks)

	// 2
	// [...] -> based on parameters passing in the initializer array will be created
	// so no need to mention fixed size of array
	kms := [...]int{10, 20, 30}
	fmt.Printf("value: %v, type: %T\n", kms, kms)

	// 3
	// empty array
	var snames [3]string
	fmt.Printf("empty array: %v\n", snames)

	// 4
	// filling empty arrays
	snames[0] = "Shiva"
	snames[2] = "Raju"
	snames[1] = "Ali"
	fmt.Printf("array: %v\n", snames)
	fmt.Printf("array[0]: %v\n", snames[0])

}

// output:

// 1
// value: [80 90 100], type: [3]int

// 2
// value: [10 20 30], type: [3]int

// 3
// empty array: [  ]

// 4
// array: [Shiva Ali Raju]
// array[0]: Shiva
