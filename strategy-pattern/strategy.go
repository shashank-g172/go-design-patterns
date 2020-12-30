package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	MarkDown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListToItem(builder *strings.Builder, item string)
}

type MarkDownListStrategy struct{}

func (m *MarkDownListStrategy) Start(builder *strings.Builder) {

}

func (m *MarkDownListStrategy) End(builder *strings.Builder) {

}

func (m *MarkDownListStrategy) AddListToItem(builder *strings.Builder, item string) {
	builder.WriteString("*" + item + "\n")
}

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListToItem(builder *strings.Builder, item string) {
	builder.WriteString("<li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case MarkDown:
		t.listStrategy = &MarkDownListStrategy{}
	case Html:
		t.listStrategy = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	s := t.listStrategy
	s.Start((&t.builder))
	for _, item := range items {
		s.AddListToItem(&t.builder, item)
	}
	s.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkDownListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	// Here, you can dynamically switch the strategy
	tp.Reset()
	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

}
