package main

import "fmt"

// Go is a strictly typed language
// i,e we have to specify value along with its type

// package level variable declaration

// In package lvl - will not be able to decalre for ex: `var1 := 3.141`
// but can be fully declare the var

var var1 float32 = 10.01
var var2 float64

func main() {

	// functional level variable declaration

	// accessing pck lvl variables
	fmt.Println("var2 = ", var1)
	fmt.Printf("%T, %v \n", var1, var1)

	var2 = 100.001
	fmt.Println("var2 = ", var2)
	fmt.Printf("%T, %v \n", var2, var2)

	// functional lvl varibales
	var3 := 1000.0001
	fmt.Println("var3 = ", var3)
	fmt.Printf("%T, %v \n", var3, var3)

}
