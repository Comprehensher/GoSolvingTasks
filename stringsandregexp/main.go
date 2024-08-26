package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(" |boat|one")
	description := "Kayak. A boat for one person."
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}
}
