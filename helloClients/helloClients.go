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
	price, error := productlist.GetProductPrice("Milk")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("Milk has price %v", price)
	message, error := productlist.GetAdvertising("Milk")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}
