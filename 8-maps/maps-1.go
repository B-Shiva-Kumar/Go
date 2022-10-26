package main

import (
	"fmt"
)

func main() {

	mapsRef()
	iterateMaps()

}

func mapsRef() {

	// 	Maps Are References:
	//  Maps are references to hash tables.

	// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

	var a = map[string]int{
		"Virat":  82,
		"DK":     0,
		"Ashwin": 1,
	}

	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["DK"] = 1
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
}

func iterateMaps() {
	// Iterating Over Maps:
	// You can use range to iterate over maps.

	var a = map[string]int{
		"Virat":  82,
		"DK":     1,
		"Ashwin": 1,
	}

	for k, v := range a {
		fmt.Printf("%v :%v, \n", k, v)
	}
}
