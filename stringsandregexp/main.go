package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "It was a boat. A small boat."
	elements := strings.Fields(text)
	joined := strings.Join(elements, "--")
	fmt.Println("Joined:", joined)
}
