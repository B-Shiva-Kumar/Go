package main

import "fmt"

func main() {

	// 2-ways to delcare arrays

	// 1
	var matrix2D [2][2]int
	matrix2D[0] = [2]int{0, 1}
	matrix2D[1] = [2]int{1, 0}
	fmt.Printf("%v\n", matrix2D)

	// 2
	// var matrix2d [2][2]int = [2][2]int{[2]int{1, 0}, [2]int{0, 1}}
	var matrix2d [2][2]int = [2][2]int{{1, 0}, {0, 1}}
	fmt.Printf("%v", matrix2d)

}

// output:

// 1
// [[0 1] [1 0]]

// 2
// [[1 0] [0 1]]
