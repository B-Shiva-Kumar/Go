package main

import "fmt"

type generic interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height, width float64
}

type triangle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * 2 * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (t triangle) area() float64 {
	return 0.5 * 3.14 * t.radius
}

func calculate_area(g generic) {
	fmt.Println(g.area())
}

func main() {
	cirlce_1 := circle{2}
	rec_1 := rectangle{2, 2}
	tri_1 := triangle{5}

	calculate_area(cirlce_1) // 12.56
	calculate_area(rec_1)    // 4
	calculate_area(tri_1)    // 7.8500000000000005
}
