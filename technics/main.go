package main

import "fmt"

func main() {
	price, tax, inStock := 275.00, 27.50, true
	fmt.Println("Total:", price+tax)
	fmt.Println("In stock:", inStock)
	price2, tax := 200.00, 25.00
	fmt.Println("Total 2:", price2+tax)
}
