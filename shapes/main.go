package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{height: 10, base: 10}
	s := square{sideLength: 7}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	result := 0.5 * t.base * t.height
	return result
}

func (s square) getArea() float64 {
	result := s.sideLength * s.sideLength
	return result
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
