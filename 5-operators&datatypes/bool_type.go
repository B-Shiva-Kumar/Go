package main

import "fmt"

func main() {
	// boolean type

	var boo bool = false
	fmt.Printf("%v %T \n", boo, boo)

	var var1 bool = 1 == 1
	fmt.Printf("%v %T\n", var1, var1)

	var2 := 2 == 1
	fmt.Printf("%v %T\n", var2, var2)

	// by default, value is 0 - false
	var var3 bool
	fmt.Printf("%v %T \n", var3, var3)

}
