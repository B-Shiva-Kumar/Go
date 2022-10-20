package main

import "fmt"

func main() {

	var (
		a int = 10 // 1010
		b int = 3  // 0011
	)

	fmt.Printf("%v \n", a&b)  //0010
	fmt.Printf("%v \n", a|b)  //1011
	fmt.Printf("%v \n", a^b)  //1001
	fmt.Printf("%v \n", a&^b) //0110

}

// output:

// 2
// 11
// 9
// 8
