// In Go, there are two functions that can be used to return the length and capacity of a slice:

// len() function - returns the length of the slice (the number of elements in the slice)
// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

// This example shows how to create slices using the []datatype{values} format:

package main

import (
	"fmt"
)

func main() {
	s1 := []int{}
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)
	fmt.Printf("%v %T \n\n", s1, s1)

	s2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2)
	fmt.Printf("%v %T \n\n", s2, s2)

	// make func
	s3 := make([]int, 5, 100)
	fmt.Println(s3)
	fmt.Printf("Length %v \n", len(s3))
	fmt.Printf("Capasity %v \n", cap(s3))

	// append (push)
	s3 = append(s3, []int{99}...) //concatenation (in JS -> spreading)
	fmt.Println(s3)
	fmt.Printf("Length %v \n", len(s3))
	fmt.Printf("Capasity %v \n", cap(s3))

	s3 = append(s3, 999, 9999)
	fmt.Println(s3)
	fmt.Printf("Length %v \n", len(s3))
	fmt.Printf("Capasity %v \n", cap(s3))

	// shift (pop)
	s4 := []int{1, 2, 3, 4, 5}

	r1 := s4[1:] // pop element from beginging
	fmt.Println(r1)

	r1 = s4[:len(s4)-1] // pop element from end
	fmt.Println(r1)

	r1 = append(s4[:2], s4[3:]...)
	fmt.Println(r1)

	// NOTE: In slices, whatever we change will reflect to original slice also
	fmt.Println(s4)

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

// [0 0 0 0 0]
// Length 5
// Capasity 100

// [0 0 0 0 0 99]
// Length 6
// Capasity 100

// [0 0 0 0 0 99 999 9999]
// Length 8
// Capasity 100

// [2 3 4 5]
// [1 2 3 4]
// [1 2 4 5]
