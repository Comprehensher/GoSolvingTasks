package main

import "fmt"

func main() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle",
		"Hat"}
	deleted := append(products[:2], products[3:]...)
	fmt.Println("Deleted:", deleted)
}
