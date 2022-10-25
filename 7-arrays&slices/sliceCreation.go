package main

import (
	"fmt"
)

func main() {

	// Create a Slice With []datatype{values}
	// Syntax:
	// slice_name := []datatype{values}
	// myslice := []int{1,2,3}

	// 1
	// Create a Slice From an Array
	// Syntax:
	// var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end] // A slice made from the array

	arr1 := [4]int{9, 99, 999, 9999}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n\n", cap(myslice))

	// 2
	// Create a Slice With The make() Function

	// Syntax:
	// slice_name := make([]type, length, capacity)
	// Note: If the capacity parameter is not defined, it will be equal to length.

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n\n", cap(myslice1))

	// 3
	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

}

// output:

// 1
// myslice = [999 9999]
// length = 2
// capacity = 2

// 2
// myslice1 = [0 0 0 0 0]
// length = 5
// capacity = 10

// 3
// myslice2 = [0 0 0 0 0]
// length = 5
// capacity = 5
