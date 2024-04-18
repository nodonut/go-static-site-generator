package main

import "fmt"

func main() {
	node := NewTextNode("Hello world", "text_type_text", "")
	props := map[string]string{
		"label": "checking",
	}
	htmlNode := NewHTMLNode("div", "i am div", props, nil)
	leafNode := NewLeafNode("p", "i am a paragrpah", props)
	parentNode := NewParentNode("div", nil, []LeafNode{leafNode})
	fmt.Println(node)
	fmt.Println(htmlNode)
	fmt.Println(leafNode.ToHtml())
	fmt.Println(parentNode.ToHtml())
}
