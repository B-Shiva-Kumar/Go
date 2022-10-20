package main

import (
	"fmt"
)

var var1 = 69 // pkg lvl

func main() {

	// Initliaze & declaring constant
	const const1 int16 = 99
	const const2 = 99.99
	const const3 = "Go gopher"
	const const4 = true

	// 1
	fmt.Printf("%v %T\n", const1, const1)
	fmt.Printf("%v %T\n", const2, const2)
	fmt.Printf("%v %T\n", const3, const3)
	fmt.Printf("%v %T\n", const4, const4)

	// 2
	// adding const & var values
	var const5 = 96 // local scope or block lvl
	fmt.Printf("%v %T\n", const5+var1, const5+var1)

}

// output:

// 1
// 99 int16
// 99.99 float64
// Go gopher string
// true bool

// 2
// 165 int
