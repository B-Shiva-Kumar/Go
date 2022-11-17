package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base, height float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func main() {

	t1 := triangle{2, 1}
	s1 := square{2}

	printArea(t1)
	printArea(s1)
}
