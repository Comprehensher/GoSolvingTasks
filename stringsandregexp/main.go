package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description,
		isLetterB))
}
