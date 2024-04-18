package main

import "fmt"

type TextNode struct {
	Text     string
	TextType string
	Url      string
}

func NewTextNode(text, textType, url string) TextNode {
	return TextNode{
		Text:     text,
		TextType: textType,
		Url:      url,
	}
}

func (t TextNode) String() string {
	return fmt.Sprintf("TextNode('%s', '%s', '%s')", t.Text, t.TextType, t.Url)
}
