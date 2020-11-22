package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions (interfaces instead of Java abstract classes, etc.)

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// other useful stuff here
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}

	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	// Parent is the parent of the child
	rs.relations = append(rs.relations,
		Info{parent, Parent, child})
	// Child is the child of the parent
	rs.relations = append(rs.relations,
		Info{child, Child, parent})
}

type Research struct {
	// relationships Relationships
	browser RelationshipBrowser // low-level
}

func (r *Research) Investigate() {
	//relations := r.relationships.relations
	//for _, rel := range relations {
	//	if rel.from.name == "John" &&
	//		rel.relationship == Parent {
	//		fmt.Println("John has a child called", rel.to.name)
	//	}
	//}

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main_5() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{&relationships}
	research.Investigate()
}
