package main

import (
	"fmt"
)

func main() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}

	for key, value := range products {
		fmt.Println("Key:", key, "Value:", value)
	}
}
