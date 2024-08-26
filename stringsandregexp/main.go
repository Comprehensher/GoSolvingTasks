package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "It was a boat. A small boat."
	replacer := strings.NewReplacer("boat", "kayak", "small",
		"huge")
	replaced := replacer.Replace(text)
	fmt.Println("Replaced:", replaced)
}
