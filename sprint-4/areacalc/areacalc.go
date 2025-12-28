package areacalc

import (
	"strings"
)

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length, width float64, shapeType string) *Rectangle {
	return &Rectangle{length: length, width: width}
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Type() string {
	return "rectangle"
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64, shapeType string) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Area() float64 {
	return pi * c.radius * c.radius
}

func (c *Circle) Type() string {
	return "circle"
}

func AreaCalculator(figures []Shape) (string, float64) {
	var totalArea float64
	shapeTypes := make([]string, 0, len(figures))

	for _, figure := range figures {
		totalArea += figure.Area()
		shapeTypes = append(shapeTypes, figure.Type())
	}

	return strings.Join(shapeTypes, "-"), totalArea
}
