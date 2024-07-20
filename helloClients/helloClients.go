package main

import (
	"fmt"
	"log"

	"eu.corp/productlist"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("From shop: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	price, error := productlist.GetProductPrice("Milk222")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("Milk has price %v", price)
}
