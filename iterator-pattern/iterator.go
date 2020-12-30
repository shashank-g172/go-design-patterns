package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

type PersonNameIterator struct {
	// Point to the actual object so we can iterate and get what we need
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("This should never be reached")
}

func main() {

	p := Person{"Alex", "Graham", "Bell"}
	for _, name := range p.Names() {
		fmt.Println(name)
	}
	l := Person{"Michael", "", "David"}
	for generatedName := range l.NamesGenerator() {
		fmt.Println(generatedName)
	}

	m := Person{"Sam", "Max", "John"}

	for it := NewPersonNameIterator(&m); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
