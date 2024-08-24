package main

import (
	"fmt"
	"strconv"
)

func main() {
	var price []rune = []rune("€48.95")
	var currency string = string(price[0])
	var amountString string = string(price[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}
