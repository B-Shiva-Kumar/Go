package main

import (
	"fmt"
)

func main() {

	// creating map function
	countryCodes := map[string]int{
		"India":      +91,
		"Afganistan": +93,
		"Algeria":    +213,
	}
	fmt.Println(countryCodes)

	// creating map using make func
	// initializing
	countryCodes1 := make(map[string]int)
	fmt.Println(countryCodes1)

	countryCodes1["India"] = +91
	countryCodes1["Afganistan"] = +93
	countryCodes1["Algeria"] = +213
	fmt.Println(countryCodes1)

	// accessing the maps using key names
	fmt.Println(countryCodes1["India"])
	fmt.Println(countryCodes1["Algeria"])

	// updating & adding the key-value pair in the map func
	countryCodes1["Russia"] = +7 // adding
	fmt.Println(countryCodes1)

	// deleting key-value from map func
	delete(countryCodes1, "Afganistan")
	fmt.Println(countryCodes1)

	// To check particular key is present in func or not
	// accessing the key-value pair which is not there in map func
	fmt.Println(countryCodes1["USA"]) // outputs Zero -> 0

	temp, ok := countryCodes1["USA"]
	fmt.Println(temp, ok) // 0 false

	// _, ok := countryCodes1["US"]
	_, ok = countryCodes1["Russia"]
	fmt.Println(ok) // true

	// len of the map func
	fmt.Println(len(countryCodes1))

	// NOTE:
	// maps are reference type
	// if any change occurs in derived ype reflecs same changes in base type also.
	fmt.Println(countryCodes)

	cp := countryCodes
	delete(countryCodes, "Afganistan")
	fmt.Println(cp)
	fmt.Println(countryCodes)

	emptyMap()

}

func emptyMap() {

	// Note: The make()function is the right way to create an empty map.
	// If you make an empty map in a different way and write to it, it will causes a runtime panic.

	// This example shows the difference between declaring an empty map using with the make()function and without it.

	// creating empty map func
	var a map[string]int
	fmt.Println("empty map without make func a : ", a)
	fmt.Println(a == nil)

	// creating empty map func with make func
	b := make(map[int]int)
	fmt.Println("empty map without make func b : ", b)
	fmt.Println(b == nil)

}
