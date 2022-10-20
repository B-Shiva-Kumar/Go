package main

import "fmt"

func main() {

	// iota: it is constructor & increamenting the value of constant
	// iota scoped to const block only

	const (
		n1 = iota
		n2 = iota
		n3 = iota
	)

	// 1
	fmt.Printf("%v %T\n", n1, n1)
	fmt.Printf("%v %T\n", n2, n2)
	fmt.Printf("%v %T\n", n3, n3)

	const (
		n4 = iota
		n5
		n6
	)

	// 2
	// if we are in same block, iota declared once & same pattern follows for others constants also
	// counter resets to Zero if we declared in the next const block
	fmt.Printf("%v %T\n", n4, n4)
	fmt.Printf("%v %T\n", n5, n5)
	fmt.Printf("%v %T\n", n6, n6)

}

// output:

// 1
// 0 int
// 1 int
// 2 int

// 2
// 0 int
// 1 int
// 2 int
