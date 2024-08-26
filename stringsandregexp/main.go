package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"
	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description,
		"o"))
	fmt.Println("IndexAny:", strings.IndexAny(description,
		"abcd"))
	fmt.Println("LastIndex:", strings.LastIndex(description,
		"o"))
	fmt.Println("LastIndexAny:",
		strings.LastIndexAny(description, "abcd"))
}
