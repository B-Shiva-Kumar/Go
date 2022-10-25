package main

import "fmt"

func main() {

	// 1
	// 	A common way of declaring a slice is like this:
	myslice := []int{}
	fmt.Printf("%v %T \n", myslice, myslice)
	// The code above declares an empty slice of 0 length and 0 capacity.

	// 2
	// To initialize the slice during declaration, use this:\
	myslice1 := []int{1, 2, 3}
	fmt.Printf("%v %T \n", myslice1, myslice1)

	// The code above declares a slice of integers of length 3 and also the capacity of 3.

	// 3
	slc1 := []int{9, 99, 999, 9999}
	fmt.Printf("%v %T \n", slc1, slc1)

}

// output:

// 3
// [9 99 999 9999] []int
