package interfacedemo

import "math"

// Shape 演示最常见的接口多态用法：定义行为，不关心具体类型。
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

// TotalArea 通过接口切片聚合不同实现类型的面积。
func TotalArea(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

