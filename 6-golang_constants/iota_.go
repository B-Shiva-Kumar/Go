package main

import "fmt"

func main() {

	// 1
	// _ -> write only variable
	// write only variable(_) is used to save the memory instead of string the value 0 as we will not use it
	// i,e. Zero value is not avaible for this _ contant
	const (
		// _ = iota  -> value 0 is not stored hence not available
		_ = iota + 3
		n1
	)

	fmt.Printf("%v %T\n", n1, n1)

	// 2
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)

}

// output:

// 1
// 4 int

// 2
// 3.73GB
