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

func main() {

	m := Person{"Miss", "Moneypenny"}
	b := SecretAgent{Person{"James", "Bond"}, true}

	m.speak()
	b.speak()

}
