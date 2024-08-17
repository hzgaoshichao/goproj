package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 5}

	fmt.Println("Rectangle")
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())

	fmt.Println("\nCircle")
	fmt.Println("Area: ", c.area())
	fmt.Println("Perimeter: ", c.perimeter())

	// 值接收者实现的结构体修改不影响接口中对应的值
	r.width = 6
	fmt.Println("\nRectangle after modification")
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())

	// 指针接收者实现的结构体修改会影响接口中对应的值
	c.radius = 7
	fmt.Println("\nCircle after modification")
	fmt.Println("Area: ", c.area())
	fmt.Println("Perimeter: ", c.perimeter())

	var gi Geometry
	gi = &r
	fmt.Println("Perimeter: ", gi.perimeter())

	gi = r
	fmt.Println("Perimeter: ", gi.perimeter())

	gi = &c
	fmt.Println("Perimeter: ", gi.perimeter())
}
