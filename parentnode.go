package main

import (
	"errors"
	"fmt"
)

type ParentNode struct {
	Tag      string
	Children []LeafNode
	Props    map[string]string
}

func NewParentNode(tag string, props map[string]string, children []LeafNode) ParentNode {
	return ParentNode{
		Children: children,
		Props:    props,
		Tag:      tag,
	}
}

func (p ParentNode) ToHtml() (string, error) {
	if p.Tag == "" {
		return "", errors.New("Invalid HTML: no tag")
	}

	if len(p.Children) == 0 {
		return "", errors.New("Invalid HTML: no children")
	}

	var childrenHtml string = fmt.Sprintf("<%s%s>", p.Tag, p.PropsToHtml())
	for _, child := range p.Children {
		childHtml, err := child.ToHtml()
		if err != nil {
			return "", err
		}
		childrenHtml = fmt.Sprintf("%s%s</%s>", childrenHtml, childHtml, p.Tag)
	}

	return childrenHtml, nil
}

func (p ParentNode) PropsToHtml() string {
	if p.Props == nil {
		return ""
	}

	var output string
	for key, value := range p.Props {
		output = fmt.Sprintf("%s %s='%s'", output, key, value)
	}

	return output
}

func (p ParentNode) String() string {
	return fmt.Sprintf("ParentNode('%s', '%s', '%+v')", p.Tag, p.PropsToHtml(), p.Children)
}

// def to_html(self):
//     if self.tag is None:
//         raise ValueError("Invalid HTML: no tag")
//
//     if self.children is None:
//         raise ValueError("Invalid HTML: no children")
//
//     children_html = ""
//     for child in self.children:
//         children_html += child.to_html()
//
//     return f"<{self.tag}{self.props_to_html()}>{children_html}</{self.tag}>"
