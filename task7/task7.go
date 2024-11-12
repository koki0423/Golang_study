package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	length() float64
}


type Rectangle struct {
	width float64
	height float64
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (r Rectangle) length() float64 {
	return 2*r.width+2*r.height
}


type Circle struct {
	radius float64
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) length() float64 {
	return 2 * math.Pi * c.radius
}


func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 7}

	printArea(r)
	printArea(c)

	printLength(r)
	printLength(c)
}

func printArea(g Geometry) {
	fmt.Println(g.area())
}

func printLength(g Geometry) {
	fmt.Println(g.length())
}
