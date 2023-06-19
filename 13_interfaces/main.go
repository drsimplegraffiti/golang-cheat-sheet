package main

import "fmt"

// Define an interface
type Shape interface {
    Area() float64 //Area is a method that returns a float64
    Perimeter() float64
}

// Implement the interface for a Rectangle
type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 { //r is a receiver of type Rectangle, Area is a method that returns a float64
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.width + 2*r.height
}

func main() {
    // Create an instance of Rectangle
    rectangle := Rectangle{width: 5, height: 3}

    // Call interface methods
    var shape Shape
    shape = rectangle
    fmt.Println("Area:", shape.Area())
    fmt.Println("Perimeter:", shape.Perimeter())
}
