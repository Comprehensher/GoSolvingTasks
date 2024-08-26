package main

import (
	"fmt"
	"unicode"
)

func main() {
	product := "Kayak"
	for _, char := range product {
		fmt.Println(string(char), "Upper case:",
			unicode.IsUpper(char))
	}
}
