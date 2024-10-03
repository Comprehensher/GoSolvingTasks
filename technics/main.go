package main

import "fmt"

func main() {
	price, tax, inStock := 275.00, 27.50, true
	fmt.Println("Total:", price+tax)
	fmt.Println("In stock:", inStock)
}
