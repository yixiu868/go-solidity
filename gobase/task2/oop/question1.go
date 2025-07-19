package oop

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func Calc() {
	circle := Circle{10}
	rectangle := Rectangle{10, 10}

	circleArea := circle.Area()
	circlePerimeter := circle.Perimeter()
	fmt.Printf("circle area is %f\n", circleArea)
	fmt.Printf("circle perimeter is %f\n", circlePerimeter)

	rectangleArea := rectangle.Area()
	rectanglePerimeter := rectangle.Perimeter()
	fmt.Printf("rectangle perimeter is %f\n", rectanglePerimeter)
	fmt.Printf("rectangle area is %f\n", rectangleArea)
}
