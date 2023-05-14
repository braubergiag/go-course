package main

import "fmt"

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	var res string
	if width >= len(text) {
		res = text
	} else if width < len(text) {
		res = text[:width] + "..."
	}

	fmt.Println(res)
}
