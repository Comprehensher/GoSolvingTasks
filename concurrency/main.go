package main

import (
	"fmt"
	// "time"
)

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	// time.Sleep(time.Second * 5)
	fmt.Println("main function complete calculation")

	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		if details, open := <-dispatchChannel; open {
			fmt.Println("Dispatch to", details.Customer, ":",
				details.Quantity,
				"x", details.Product.Name)
		} else {
			fmt.Println("Channel has been closed")
			break
		}
	}
}
