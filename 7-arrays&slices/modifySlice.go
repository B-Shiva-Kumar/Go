// Go Access, Change, Append and Copy Slices:

package main

import (
	"fmt"
)

func main() {

	// 1. Access Elements of a Slice:
	// You can access a specific slice element by referring to the index number.
	// In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// 2. Change Elements of a Slice
	// You can also change a specific slice element by referring to the index number.

	prices1 := []int{10, 20, 30}
	prices1[2] = 50
	fmt.Println(prices1[0])
	fmt.Println(prices1[2])

	// 3. Append Elements To a Slice
	// You can append elements to the end of a slice using the append() function.
	// Syntax: slice_name = append(slice_name, element1, element2, ...)

	myslice1 := []int{1, 2, 3, 4, 5, 6} // [1 2 3 4 5 6]
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21) // [1 2 3 4 5 6 20 21]
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// 4. Append One Slice To Another Slice
	// To append all the elements of one slice to another slice, use the append() function.
	// Syntax: slice3 = append(slice1, slice2...)
	// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := append(slice1, slice2...)

	fmt.Printf("myslice3=%v\n", slice3)
	fmt.Printf("length=%d\n", len(slice3))
	fmt.Printf("capacity=%d\n", cap(slice3))

	// 5. Change The Length of a Slice
	// Unlike arrays, it is possible to change the length of a slice.

	arr1 := [6]int{9, 99, 999, 9999, 99999, 999999} // An array
	mySlice1 := arr1[1:5]                           // Slice array
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n\n", cap(mySlice1))

	mySlice1 = arr1[2:4] // Change length by re-slicing the array
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n\n", cap(mySlice1))

	mySlice1 = append(mySlice1, 9999999, 99999999) // Change length by appending items
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n\n", cap(mySlice1))

}

// output:

// 1.
// 10
// 30

// 2.
// 10
// 50

// 3.
// myslice1 = [1 2 3 4 5 6]
// length = 6
// capacity = 6
// myslice1 = [1 2 3 4 5 6 20 21]
// length = 8
// capacity = 12

// 4.
// myslice3=[1 2 3 4 5 6]
// length=6
// capacity=6
