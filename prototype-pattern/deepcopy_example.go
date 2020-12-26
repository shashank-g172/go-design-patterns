package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopyUsingSerialization() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))
	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	// This is because it's the easiest way of copying slices
	copy(q.Friends, p.Friends)
	return &q
}

func main() {

	john := Person{"John", &Address{"123 Boston Road", "Boston", "US"}, []string{"Chris", "Matt"}}

	jane := john
	// Address gets copied as well

	jane.Name = "Jane"

	jane2 := john
	// Deep copying,  make copies of everything it refers to. This isn't a great way to do it though. This is the problem being solved here
	jane2.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country}
	jane2.Name = "Jane 2"
	jane2.Address.StreetAddress = "213 Boston Road"

	jane3 := john.DeepCopy()
	jane3.Name = "Jane 3"
	jane3.Address.StreetAddress = "400 Boston Road"
	jane3.Friends = append(jane3.Friends, "Angela")

	// Prototyping is literally taking an object, and making modifications as needed to it.
	jane4 := john.DeepCopyUsingSerialization()
	jane4.Name = "Jane 4"
	jane4.Address.StreetAddress = "500 Boston Road"
	jane4.Friends = append(jane4.Friends, "Michael")

	fmt.Println(john, john.Address, john.Friends)
	fmt.Println(jane, jane.Address)
	fmt.Println(jane2, jane2.Address)
	fmt.Println(jane3, jane3.Address, jane3.Friends)
	fmt.Println(jane4, jane4.Address, jane4.Friends)

}
