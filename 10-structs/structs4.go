package main

import (
	"fmt"
)

type vertex struct {
	x, y int
}

func main() {
	var v = vertex{1, 2}
	fmt.Println(v)

	v = vertex{x: 4, y: 6}
	fmt.Println(v)

	var w = []vertex{{1, 2}, {4, 5}, {10, 11}}
	fmt.Println(w)
}
