package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f",
		c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// possible, but not generic enough
type ColoredSquare struct {
	Square
	Color string
}

// Instead, use something more generic like this
type ColoredShape struct {
	Shape Shape
	Color string
}

// ColoredShape is also a shape, so it has to implement the Shape interface
func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s",
		c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency",
		t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())
	// cirlce.Resize(2) can be called on the circle struct directly

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())
	// redCircle.Resize() cannot be called, since ReSize is only applicable for the Circle struct
	// and not the decorator over it. This is a real life limitation of the decorator pattern

	// Decorator over Decorator is permitted - they can be composed
	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
