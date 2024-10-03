package main

import "fmt"

func main() {
	price, tax, inStock, _ := 275.00, 27.50, true, true
	var _ = "Alice"
	fmt.Println("Total:", price+tax)
	fmt.Println("In stock:", inStock)
}
