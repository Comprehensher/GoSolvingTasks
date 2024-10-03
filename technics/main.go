package main

import "fmt"

func main() {
	const price, tax float32 = 275, 27.50
	fmt.Println("price:", price)
	fmt.Println("tax:", tax)

	const CONSUME, isTrue = 1, true
	fmt.Println("CONSUME:", CONSUME)
	fmt.Println("isTrue:", isTrue)
}
