package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:",
		calculator(price))
}

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}
}
