// In Go, there are two functions that can be used to return the length and capacity of a slice:

// len() function - returns the length of the slice (the number of elements in the slice)
// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

// This example shows how to create slices using the []datatype{values} format:

package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)
	fmt.Printf("%v %T \n\n", myslice1, myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	fmt.Printf("%v %T", myslice2, myslice2)

}

// output

// 0
// 0
// []
// [] []int

// 4
// 4
// [Go Slices Are Powerful]
// [Go Slices Are Powerful] []string
