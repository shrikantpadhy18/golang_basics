package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumference() float64
}

type circle struct {
	radius float64
}

type square struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) area() float64 {
	return s.length * s.breadth
}

func (s square) circumference() float64 {
	return 2 * (s.length + s.breadth)
}

func getInfo(obj shape) {
	fmt.Printf("Area: %f\n", obj.area())
	fmt.Printf("Circumference: %f\n", obj.circumference())
}

func main() {
	c := circle{
		radius: 25.60,
	}
	s := square{
		length:  23,
		breadth: 20,
	}

	fmt.Println("Circle:")
	getInfo(c)
	fmt.Println("Square:")
	getInfo(s)
}
