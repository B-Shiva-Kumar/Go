package main

import "fmt"

func main() {

	// 1
	// copying arr1 to arr2
	// if any modification is made for arr2 will not change in arr1
	// to effect modification for both arrays -> we can use pointers
	arr1 := [...]int{1, 2, 3}
	arr2 := arr1

	arr2[0] = 100

	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n\n", arr2)

	// 2
	// using pointers -> modifing in arr4 will effects in arr3 also
	// using pointers we are not making any copy
	arr3 := [...]int{100, 1000, 1000}
	arr4 := &arr3

	fmt.Printf("%v\n", arr3)
	fmt.Printf("%v\n\n", arr4)

	arr4[0] = 99
	arr4[1] = 999
	fmt.Printf("%v\n", arr3)
	fmt.Printf("%v\n", arr4)

}

// output:

// 1
// [1 2 3]
// [100 2 3]

// 2
// [100 1000 1000]
// &[100 1000 1000]

// [99 999 1000]
// &[99 999 1000]
