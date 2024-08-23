package main

import (
	"fmt"
	"sort"
)

func main() {
	products := []string{"Kayak", "Lifejacket", "Paddle",
		"Hat"}
	sort.Strings(products)
	for index, value := range products {
		fmt.Println("Index:", index, "Value:", value)
	}
}
