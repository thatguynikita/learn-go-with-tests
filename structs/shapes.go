package main

import "math"

// Shape is the interface for all sorts of geometrical shapes
type Shape interface {
	Area() float64
}

// Rectangle is rectangle
type Rectangle struct {
	Width float64
	Height float64
}

// Area return area for the Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is circle
type Circle struct {
	Radius float64
}

// Area returns area for the Circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Triangle is triangle
type Triangle struct {
	Base float64
	Height float64
}

// Area returns area for the Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter returns the perimeter of a given rectangle
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

// Area returns the area of a given rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}