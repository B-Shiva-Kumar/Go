package main

import "fmt"

func main() {

	slc1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // slice
	fmt.Println(slc1)

	// slicing
	a := slc1[:]
	b := slc1[:6]
	c := slc1[6:]
	d := slc1[3:6]
	e := slc1[1:3]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// indexing
	slc1[1] = 99
	slc1[2] = 999
	fmt.Println(slc1)

}
