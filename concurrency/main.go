package main

import (
	"fmt"
	"time"
	// "time"
)

// func receiveDispatches(channel <-chan DispatchNotification) {
// 	for details := range channel {
// 		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
// 	}
// 	fmt.Println("Channel has been closed")
// }

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	// time.Sleep(time.Second * 5)
	fmt.Println("main function complete calculation")

	dispatchChannel := make(chan DispatchNotification, 100)
	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	// receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				goto alldone
			}
		default:
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")
}
