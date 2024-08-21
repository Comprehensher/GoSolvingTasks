package main

import "fmt"

func main() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle",
		"Hat"}
	someNames := products[1:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:",
		cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:",
		cap(allNames))
}
