package main

import "fmt"

func main() {
	node := newTextNode("Hello world", "text_type_text", "")
	fmt.Println(node)
}
