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

	productPrices, error := productlist.GetProductsPrice([]string{"Milk", "Bread"})
	if error != nil {
		log.Fatal(error)
	}
	for productName, productPrice := range productPrices {
		fmt.Printf("%v has price %v", productName, productPrice)
	}

	productInfo, error := productlist.GetProductInfo("Bread")
	if !productInfo.IsEmpty() {
		fmt.Printf(" Product name - %v ", productInfo.Name)
		fmt.Printf(" Product price - %v ", productInfo.Price)
		fmt.Printf(" Product date - %v ", productInfo.ArrivingDate.Format("2006-1-2"))
	}
	if error != nil {
		log.Fatal(error)
	}

}
