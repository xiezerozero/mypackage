package main

import "fmt"

type ShapeInterface interface {
	Area() float64
	GetName() string
	SetArea(float64)
}

// 标准形状，它的面积为0.0
type Shape struct {
	name string
	area float64
}

func (s *Shape) Area() float64 {
	return 0.0
}

func (s *Shape) GetName() string {
	return s.name
}

func (s *Shape) SetArea(area float64) {
	s.area = area
}

func (s *Shape) PrintArea() {
	fmt.Printf("%s : Area %v\r\n", s.name, s.area)
}

type Rectangle struct {
	Shape
	w, h float64
}

func (r *Rectangle) Area() float64 {
	return r.w * r.h
}

func main() {
	r := &Rectangle{
		Shape: Shape{name: "rectangle"},
		w:     5,
		h:     4,
	}
	initShape(r)
	r.PrintArea()
}

func initShape(s ShapeInterface) {
	area := s.Area()
	s.SetArea(area)
}
