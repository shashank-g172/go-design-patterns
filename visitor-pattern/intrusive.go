package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	// This is a violation of the Open closed principle. Print shouldn't be added to the Expression interface to traverse it.
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.left.Print(sb)
	sb.WriteRune('+')
	a.right.Print(sb)
	sb.WriteRune(')')
}

func main() {
	// 										1 						+ 							(2								+3)
	e := &AdditionExpression{left: &DoubleExpression{1}, right: &AdditionExpression{left: &DoubleExpression{2}, right: &DoubleExpression{3}}}
	sb := strings.Builder{}
	e.Print(&sb)
	fmt.Println(sb.String())

}
