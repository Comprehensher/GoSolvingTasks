package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace:", replace)
	fmt.Println("Replace All:", replaceAll)
}
