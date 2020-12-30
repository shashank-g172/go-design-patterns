package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

type Observer interface {
	Notify(data interface{})
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{Observable: Observable{new(list.List)}, Name: name}
}

func (p *Person) CatchCold() {
	p.Fire(p.Name)
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s", data.(string))
}

func main() {

	p := NewPerson("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	p.CatchCold()

}
