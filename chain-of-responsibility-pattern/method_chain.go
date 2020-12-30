package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifer struct {
	CreatureModifier
}

func NewDoubleAttackModifer(c *Creature) *DoubleAttackModifer {
	return &DoubleAttackModifer{CreatureModifier{creature: c}}
}

type IncreaseDefenseModifer struct {
	CreatureModifier
}

func NewIncreaseDefenseModifer(c *Creature) *IncreaseDefenseModifer {
	return &IncreaseDefenseModifer{CreatureModifier{creature: c}}
}

type NoBonusModifer struct {
	CreatureModifier
}

func NewNoBonusModifer(c *Creature) *NoBonusModifer {
	return &NoBonusModifer{CreatureModifier{creature: c}}
}

func (n *NoBonusModifer) Handle() {
	//empty
}

func (d *DoubleAttackModifer) Handle() {
	fmt.Println("Doubling \b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

func (i *IncreaseDefenseModifer) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("Increasing", i.creature.Name, "\b's defense")
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}

func main() {

	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	// If we don't want to propagate the chain, simply don't implement Handle
	// root.Add(NewNoBonusModifer(goblin))

	// Propagate the chain through the modifers
	root.Add(NewDoubleAttackModifer(goblin))
	root.Add(NewIncreaseDefenseModifer(goblin))
	root.Add(NewDoubleAttackModifer(goblin))

	root.Handle()
	fmt.Println(goblin.String())
}
