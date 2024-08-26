package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"
	prefixTrimmed := strings.TrimPrefix(description, "A boat")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	fmt.Println("Trimmed:", prefixTrimmed)
	fmt.Println("Not trimmed:", wrongPrefix)
}
