package main

import (
	"errors"
	"fmt"
)

type LeafNode struct {
	Tag   string
	Value string
	Props map[string]string
}

func NewLeafNode(tag string, value string, props map[string]string) LeafNode {
	return LeafNode{
		Tag:   tag,
		Value: value,
		Props: props,
	}
}

func (l LeafNode) ToHtml() (string, error) {
	if l.Value == "" {
		return "", errors.New("value is required")
	}

	if l.Tag == "" {
		return l.Value, nil
	}

	return fmt.Sprintf("<%s%s>%s</%s>", l.Tag, l.PropsToHtml(), l.Value, l.Tag), nil
}

func (l LeafNode) PropsToHtml() string {
	if l.Props == nil {
		return ""
	}

	var output string
	for key, value := range l.Props {
		output = fmt.Sprintf("%s %s='%s'", output, key, value)
	}

	return output
}

func (l LeafNode) String() string {
	return fmt.Sprintf("LeafNode('%s', '%s', '%+v')", l.Tag, l.Value, l.PropsToHtml())
}
