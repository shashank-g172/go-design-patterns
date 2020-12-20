package main

import "fmt"

type Person interface {
	SayHello()
}

type tiredPerson struct {
	name string
	age  int
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s and I am %d years old\n", p.name, p.age)
}

func (t *tiredPerson) SayHello() {
	fmt.Println("Can't talk, too tired")
}

func NewPerson(name string, age int) Person {
	// Hide the logic of what person is being created underneath the interface instantiation, instead
	// of giving a user control of it
	if age > 80 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("TestUser", 90)
	p.SayHello()
}
