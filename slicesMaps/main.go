package main

import "fmt"

func main() {
	products := []string{"Kayak", "Lifejacket", "Paddle",
		"Hat"}
	for index, value := range products[2:] {
		fmt.Println("Index:", index, "Value:", value)
	}
}
