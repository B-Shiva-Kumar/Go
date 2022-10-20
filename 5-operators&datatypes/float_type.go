package main

import "fmt"

func main() {

	var (
		x   float32 = 9.99e20 //9.99x10^20
		y   float64 = 3.14
		sum float32
	)

	// sum = x + y
	// x + y (mismatched types float32 and float64)

	// 1
	sum = x + float32(y)
	fmt.Printf("x - %v & %T\n", x, x)
	fmt.Printf("y - %v & %T\n", y, y)
	fmt.Printf("sum value: %v\n", sum)
	fmt.Printf("sum type: %T\n", sum)

	// NOTE:  The default type for float is float64. If you do not specify a type, the type will be float64.
	var (
		a = 3.14
		b = 10.12e10
	)

	// 2
	sum1 := a + b
	fmt.Printf("a - %v & %T\n", a, a)
	fmt.Printf("b - %v & %T\n", b, b)
	fmt.Printf("sum1 value: %v\n", sum1)
	fmt.Printf("sum1 type: %T\n", sum1)

	// 3
	// This example will result in an error because 3.4e+39 is out of range for float32:
	// var z float32 = 3.4e+39
	// fmt.Println(z)

	// 4
	// arithmetic operators with floating pnt numbers
	// remainder operator(%), right sift(>>), left shift(<<) won't be able to use with float data types
	fmt.Printf("%v %T \n", a+b, a+b)
	fmt.Printf("%v %T \n", a+b, a-b)
	fmt.Printf("%v %T \n", a+b, a*b)
	fmt.Printf("%v %T \n", a+b, a/b)

}

// output:

// 1
// x - 9.99e+20 & float32
// y - 3.14 & float64
// sum value: 9.99e+20
// sum type: float32

// 2
// a - 3.14 & float64
// b - 1.012e+11 & float64
// sum1 value: 1.0120000000314e+11
// sum1 type: float64

// 3
// cannot use 3.4e+39 (untyped float constant) as float32 value in variable declaration (overflows)

// 4
// 1.0120000000314e+11 float64
// 1.0120000000314e+11 float64
// 1.0120000000314e+11 float64
// 1.0120000000314e+11 float64
