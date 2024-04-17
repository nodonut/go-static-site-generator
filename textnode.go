package main

import "fmt"

type textNode struct {
	Text     string
	TextType string
	Url      string
}

type TextNode interface {
}

func newTextNode(text, textType, url string) textNode {
	return textNode{
		Text:     text,
		TextType: textType,
		Url:      url,
	}
}

func (t textNode) String() string {
	return fmt.Sprintf("TextNode('%s', '%s', '%s')", t.Text, t.TextType, t.Url)
}
