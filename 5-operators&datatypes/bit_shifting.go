package main

import "fmt"

func main() {

	// NOTE:
	// for right shift -> mul (*)
	// for left shift  -> div (/)

	var (
		a int = 8      // 1000
		b     = a << 3 // (2^3 * 2^3)
		c     = a >> 3 // (2^3 / 2^3)
	)

	fmt.Println("left shift:", b)
	fmt.Println("right shift:", c)

	var (
		x int8 = 10          // 1010
		y int  = int(x) << 2 // [(2^2 + 2^0) * 2^2]
		z      = x >> 2      // [(2^2 + 2^0) / 2^2]
	)

	fmt.Println("left shift:", y)
	fmt.Println("right shift:", z)

}

// output:

// left shift: 64
// right shift: 1

// left shift: 40
// right shift: 2
