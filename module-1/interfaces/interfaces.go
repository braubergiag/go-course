package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func measure(g geometry) {
	fmt.Printf("%T: %+v\n", g, g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimiter", g.perim())
}
func main() {
	r := rect{width: 3, height: 3}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
