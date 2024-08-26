package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "It was a boat. A small boat."

	var builder strings.Builder
	for _, sub := range strings.Fields(text) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	fmt.Println("String:", builder.String())
}
