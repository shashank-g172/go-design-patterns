package main

import "fmt"

type Buffer struct {
	width, height int
	bufer         []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height, make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.bufer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewPort(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

type Console struct {
	buffer    []*Buffer
	viewPorts []*Viewport
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewPort(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewPorts[0].GetCharacterAt(index)
}
func main() {
	// Instead of working with buffers and viewports manually, just work with the Console Facade
	// which takes care of initialization underneath
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
