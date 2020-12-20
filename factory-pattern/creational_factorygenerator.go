package main

import (
	"fmt"
)

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

//functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

//Structural
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactoryStructural(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {

	fmt.Printf("First way, functional")
	developerFactory := NewEmployeeFactory("developer", 160000)
	managerFactory := NewEmployeeFactory("manager", 280000)

	// Simply fill in the missing information to create the person
	developer := developerFactory("NewGuy")
	manager := managerFactory("notNewguy")

	fmt.Println(developer)
	fmt.Println(manager)

	fmt.Printf("Second way, structural. Provides a way to change values after instantiation")

	developerFactory2 := NewEmployeeFactoryStructural("developer", 160000)
	developerFactory2.AnnualIncome = 180000
	managerFactory2 := NewEmployeeFactoryStructural("manager", 280000)

	developer2 := developerFactory2.Create("NewGuy")

	manager2 := managerFactory2.Create("notNewguy")

	fmt.Println(developer2)
	fmt.Println(manager2)

}
