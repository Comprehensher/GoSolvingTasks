package main

import (
	"fmt"
	"time"
)

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel <- p:
			fmt.Println("Sent product:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("main function complete calculation")

	productChannel := make(chan *Product, 5)
	go enumerateProducts(productChannel)

	time.Sleep(time.Second)

	for p := range productChannel {
		fmt.Println("Received product:", p.Name)
	}
}
