package main

import (
	"fmt"

	"eu.corp/productlist"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Printf("Milk has price %v", productlist.GetProductPrice("Milk"))
}
