package main

import (
	//"fmt"
	"time"
)

func writeToChannel(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 3)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	channelOpen := true
	for channelOpen {
		Printfln("Starting channel read")
		select {
		case name, ok := <-nameChannel:
			if !ok {
				channelOpen = false
				break
			} else {
				Printfln("Read name: %v", name)
			}
		case <-time.After(time.Second * 2):
			Printfln("Timeout")
		}
	}
}
