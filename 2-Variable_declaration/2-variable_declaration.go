package main

import "fmt"

func main() {

	// once variable is declared we must use it
	// else throws error

	// we can declare variables in 3-ways

	// 1st
	var var1 int // variable declaration
	var1 = 100   // variable initialization

	fmt.Println("var1 = ", var1)
	// type of variable <- %T & value <- %v
	fmt.Printf("%T, %v\n", var1, var1)

	// 2nd
	var var2 int = 1000
	fmt.Println("var2 = ", var2)
	fmt.Printf("%T, %v \n", var2, var2)

	// 3rd
	// short hand operator
	var3 := 10000
	fmt.Println("var3 = ", var3)
	fmt.Printf("%T %v", var3, var3)

}
