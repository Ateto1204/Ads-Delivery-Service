package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{
		width:  10,
		height: 5,
	}

	c := circle{
		radius: 4,
	}

	fmt.Println("Area of rectangle: ", r.area())
	fmt.Println("Perimeter of rectangle: ", r.perimeter())

	fmt.Println("Area of circle: ", c.area())
	fmt.Println("Perimeter of circle: ", c.perimeter())
}
