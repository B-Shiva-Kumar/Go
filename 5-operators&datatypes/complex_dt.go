package main

import "fmt"

func main() {

	var (
		x complex64  = 2 + 2i
		y complex128 = 4 + 4i
		z complex128 = complex(99.9e99, 99.9e99)
	)

	// 1
	fmt.Printf("x - values: %v & type: %T\n", x, x)
	fmt.Printf("y - values: %v & type: %T\n", y, y)
	fmt.Printf("z - values: %v & type: %T\n", z, z)

	// 2
	// arithmetic operations with complex data type
	var (
		a = complex(2, 2)
		b = complex(22, 22)
	)
	fmt.Printf("%v %T \n", a+b, a+b)
	fmt.Printf("%v %T \n", a-b, a-b)
	fmt.Printf("%v %T \n", a*b, a*b)
	fmt.Printf("%v %T \n", a/b, a/b)

	// 3
	// extracting real & imag part from the complex value
	// using built-in functions real() & imag()
	fmt.Printf("real part of a - value: %v & type: %T \n", real(a), real(a))
	fmt.Printf("imag part of a - value: %v & type: %T \n", imag(a), imag(a))

}

// output:

// 1
// x - values: (2+2i) & type: complex64
// y - values: (4+4i) & type: complex128
// z - values: (9.99e+100+9.99e+100i) & type: complex128

// 2
// (24+24i) complex128
// (-20-20i) complex128
// (0+88i) complex128
// (0.09090909090909091+0i) complex128

// 3
// real part of a - value: 2 & type: float64
// imag part of a - value: 2 & type: float64
