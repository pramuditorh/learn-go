package main

import (
	"fmt"
	"math"
	"reflect"
)

// Circle struct is exported (upper case 'C') and accessible for other packages.
type Circle struct {
	radius float64
}

// rectangle struct is not exported (lower case 'r') and not accessible for other packages.
type rectangle struct {
	length, width float64
}

// Shape (declare Shape interface) interface is exported
type Shape interface {
	area() float64
}

// MultiShape struct use interface Shape as a array field
type MultiShape struct {
	shapes []Shape
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.radius * c.radius
}

// Use methods
func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *rectangle) area() float64 {
	return r.length * r.width
}

func totalArea(shapes ...Shape) float64 {
	var totalArea float64

	for _, s := range shapes {
		totalArea += s.area()
	}

	return totalArea
}

// return Circle area or rectangle area
func myArea(s Shape) float64 {
	return s.area()
}

func (m *MultiShape) area() float64 {
	var totalArea float64

	for _, s := range m.shapes {
		totalArea += s.area()
	}

	return totalArea
}

func main() {
	c := Circle{radius: 5}
	r := rectangle{10, 7}
	// Init MultiShape struct
	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{14},
			&rectangle{20, 12},
		},
	}

	/*
		Other ways to init struct
		var c Circle
		c1 := Circle(0, 0, 5)
		c1 := &Circle(0, 0, 5)
	*/

	fmt.Println(c)
	fmt.Println(c.radius)
	fmt.Println(reflect.TypeOf(c))        // print main.Circle
	fmt.Println(reflect.TypeOf(c.radius)) // print float64
	fmt.Println("print multiShape: ", multiShape.shapes[0])

	// Struct are mutable
	c.radius = 3
	fmt.Println(c.radius)

	fmt.Println("circleArea = ", circleArea(&c))
	fmt.Println("c.area() = ", c.area())
	fmt.Println("r.area() = ", r.area())
	fmt.Println("totalArea interface = ", totalArea(&c, &r))
	fmt.Println("Circle area from totalArea = ", totalArea(&c))
	fmt.Println("myArea interface (Circle area) = ", myArea(&c))
	fmt.Println("MultiShape struct = ", multiShape.area())
	fmt.Println("MultiShape struct (rectangle area) = ", multiShape.shapes[1].area())
}
