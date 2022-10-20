package main

import "fmt"

func main() {

	var (
		s1 string = "Hello"
		s2 string = "World"
	)

	// 1
	fmt.Printf("s1 - value: %v & type: %T\n", s1, s1)
	fmt.Printf("s2 - value: %v & type: %T\n", s2, s2)

	// 2
	// concatination
	s3 := s1 + s2
	fmt.Printf("s3 - value: %v & type: %T\n", s3, s3)

	// 3
	// indexing
	// without conversion using string -> we will get asscii value & its type
	fmt.Printf("s1 - value: %v, value: %v & type: %T, type: %T\n", s1[0], string(s1[0]), s1[0], string(s1[0]))

	// 4
	// slice of bytes/ array of bytes/ collection of bytes
	s4 := []byte(s1)
	fmt.Printf("s4 - value: %v & type: %T\n", s4, s4)

}

// output:

// 1
// s1 - value: Hello & type: string
// s2 - value: World & type: string

// 2
// s3 - value: HelloWorld & type: string

// 3
// s1 - value: 72, value: H & type: uint8, type: string

// 4
// s4 - value: [72 101 108 108 111] & type: []uint8
