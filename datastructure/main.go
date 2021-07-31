package main

import "fmt"
type Vertex struct {
	X int
	Y int
  }
  type Shape interface {
	Area() float64
	Perimeter() float64
  }
  type Rectangle struct {
	Length, Width float64
  }
  func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)
	var r Shape = Rectangle{Length: 3, Width: 4}
  fmt.Printf("Type of r: %T, Area: %v, Perimeter: %v.", r, r.Area(), r.Perimeter())
  }

  func (r Rectangle) Area() float64 {
	return r.Length * r.Width
  }
  
  func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
  }