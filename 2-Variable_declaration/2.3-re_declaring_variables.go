package main

import "fmt"

// pck lvl variables
var (
	a int     = 1
	b float32 = 999.9
)

func main() {

	// NOTE:
	// 1. we can re assign values in the local scope when already declared in global scope
	// 2. we can initialize & declare the existing variable in the local scope (functional block) when already declared in the global scope (pck lvl)
	// 3. we cannot re-declare(existing var) a variable in the same block

	// Shadowing Principle:
	// Shadowing is a principle when a variable overrides a variable in a more specific scope.
	// This means that when a variable is declared in an inner scope having the same data type and name in the outer scope, the variable is said to be `shadowed`.

	// 1
	// variable re-assigning value
	fmt.Println("a = ", a)
	a = 2
	fmt.Println("a = ", a)

	// variable Initialization & re-declaration of the existing variable
	var a string = "Golang"
	fmt.Println("a = ", a)

	// NOTE: we will not able to re-declare variable in the functional local scope when Initialized & Declared already
	// a := 22
	// fmt.Println("a = ", a)

	// 2
	// re-assign pcl lvl variable here with different type
	fmt.Println("b = ", b)
	b = 999.999
	fmt.Println("b = ", b)

	// var b string = "Go"
	c := "Go gopher"
	fmt.Println("c = ", c)

	// 3
	// var d string = "1st"
	// fmt.Println(d)

	// var d int = 2
	// fmt.Println(d)

}
