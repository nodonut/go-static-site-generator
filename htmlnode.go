package main

import "fmt"

type HTMLNode struct {
	Tag      string
	Value    string
	Children []HTMLNode
	Props    map[string]string
}

func NewHTMLNode(tag string, value string, props map[string]string, children []HTMLNode) HTMLNode {
	return HTMLNode{
		Children: children,
		Props:    props,
		Tag:      tag,
		Value:    value,
	}
}

func (h HTMLNode) PropsToHtml() string {
	if h.Props == nil {
		return ""
	}

	var output string
	for key, value := range h.Props {
		output = fmt.Sprintf("%s %s='%s'", output, key, value)
	}

	return output
}

func (h HTMLNode) String() string {
	return fmt.Sprintf("HTMLNode('%s', '%s', '%s', '%+v')", h.Tag, h.Value, h.PropsToHtml(), h.Children)
}
