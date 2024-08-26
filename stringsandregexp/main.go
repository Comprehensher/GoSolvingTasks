package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "It was a boat. A small boat."
	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}
	mapped := strings.Map(mapper, text)
	fmt.Println("Mapped:", mapped)
}
