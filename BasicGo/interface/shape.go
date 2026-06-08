package interfacedemo

import "math"

type Shape interface {
	Area() float64
}

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func TotalArea(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}
