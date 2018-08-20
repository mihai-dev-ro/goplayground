package main

import (
	"fmt"
	"math"
)

func main() {
	x := Circle{5, 5, 10}
	y := Rectangle{4, 7, 10, 4}

	fmt.Printf(
		"Circle: %v, area: %v, perimeter: %v\n", 
		x, x.area(), x.perimeter())
	fmt.Printf(
		"Rectangle: %v, width: %v, height: %v, area: %v, perimeter: %v\n", 
		y, y.width(), y.height(), y.area(), y.perimeter())

	totalArea := calcAreaTotal(&x, &y)
	fmt.Printf("Total area: %v\n", totalArea)
	totalPerimeter := calcPerimeterTotal(&x, &y)
	fmt.Printf("Total perimeter: %v\n", totalPerimeter)
}

type Shape interface {
	area() float64
	perimeter() float64
}

func calcAreaTotal(shapes ...Shape) float64 {
	total := 0.0 // force the type float64
	for _,s := range shapes {
		total += s.area()
	}
	return total
}

func calcPerimeterTotal(shapes ...Shape) float64 {
	total := 0.0 //force the type float
	for _,s := range shapes {
		total += s.perimeter()
	}
	return total
}

type Circle struct {
	X, Y, R float64
}
func (c *Circle) area() float64 {
	return math.Pi * c.R * c.R
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.R
}

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

func (r *Rectangle) width() float64 {
	if r.X1 > r.X2 {
		return (r.X1 - r.X2)
	} else {
		return (r.X2 - r.X1)
	}
}

func (r *Rectangle) height() float64 {
	if r.Y1 > r.Y2 {
		return (r.Y1 - r.Y2)
	} else {
		return (r.Y2 - r.Y1)
	}
}

func (r *Rectangle) area() float64 {
	return r.width() * r.height()
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.width() + r.height())
}