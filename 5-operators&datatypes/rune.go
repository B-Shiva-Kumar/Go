// Simple Go program to illustrate
// how to create a rune
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Creating a rune
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// Displaying rune and its type
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s; Type: %T", rune1,
		rune1, reflect.TypeOf(rune1), rune1)

	fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s; Type: %T", rune2,
		rune2, reflect.TypeOf(rune2), rune2)

	fmt.Printf("\nRune 3: Unicode: %U; Type: %s; Type: %T", rune3,
		reflect.TypeOf(rune3), rune3)

}
