package main

import "fmt"

type Person struct {
	fname, lname string
}

type SecretAgent struct {
	Person
	licenseToKill bool
}

func (p Person) speak() {
	fmt.Println(p.fname, p.lname, "says Good morning")
}

func (s SecretAgent) speak() {
	fmt.Println(s.Person.fname, s.Person.lname, "says Shaken, not stirred")
}

type Human interface {
	// This implies that all the structs (Person and SecretAgent) implicitly implement the Human interface
	// That also means that they are "of type" Human
	speak()
}

func SaySomething(h Human) {
	h.speak()
}

func main() {

	m := Person{"Miss", "Moneypenny"}
	b := SecretAgent{Person{"James", "Bond"}, true}

	SaySomething(m)
	SaySomething(b)

}
