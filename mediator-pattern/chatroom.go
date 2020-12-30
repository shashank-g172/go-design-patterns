package main

import "fmt"

type Person struct {
	Name    string
	Room    *Chatroom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s\n", sender, message)
	fmt.Printf("[%s's chat session]: %s", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type Chatroom struct {
	people []*Person
}

func (c *Chatroom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (c *Chatroom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func (c *Chatroom) Join(p *Person) {
	joinMsg := p.Name + " joined the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := Chatroom{}

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("Hi everybody")
	jane.Say("Hi John!")

	simon := NewPerson("Simon")
	room.Join(simon)

	simon.Say("Hi everyone")

	jane.PrivateMessage("Simon", "Glad you could join us")
}
