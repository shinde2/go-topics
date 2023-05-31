package main

import "fmt"

/*
  - when a type defines all methods inside the interface,
    a type is said to implement that interface
  - Embedding interfaces is allowed
  - a type can implement multiple interfaces
*/

// example -  https://gobyexample.com/interfaces

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 3, height: 4}
	var g geometry = r
	fmt.Println(g.area(), g.perim())
}
