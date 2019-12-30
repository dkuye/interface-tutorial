package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height // A = ab
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height) // u = 2(a+b)
}

func (c circle) area() float64 {
	return math.Pi * (math.Pow(c.radius, 2)) // A = πr^2
}

func (c circle) perimeter() float64 {
	return 2 * (math.Pi * c.radius) // u = 2πr
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
	fmt.Println("")
}

func main() {
	r := rect{width: 12, height: 24}
	c := circle{radius: 30}

	measure(r)
	measure(c)
}
