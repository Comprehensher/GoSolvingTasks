package main

import "fmt"

func main() {
	var price, tax = 275.00, 27.50
	fmt.Println(price + tax)

	var price1, tax1 float64
	price1 = 275.00
	tax1 = 27.50
	fmt.Println(price1 + tax1)
}
