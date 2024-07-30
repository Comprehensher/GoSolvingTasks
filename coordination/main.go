package main

import (
	"sync"
	"time"
)

func processRequest(wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		Printfln("Processing request: %v", total)
		total++
		time.Sleep(time.Millisecond * 250)
	}
	Printfln("Request processed...%v", total)
	wg.Done()
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	go processRequest(&waitGroup, 10)
	waitGroup.Wait()
}
