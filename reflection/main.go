package main

import (
	"reflect"
)

func sendOverChannel(channel interface{}, data interface{}) {
	channelVal := reflect.ValueOf(channel)
	dataVal := reflect.ValueOf(data)
	if channelVal.Kind() == reflect.Chan &&
		dataVal.Kind() == reflect.Slice &&
		channelVal.Type().Elem() ==
			dataVal.Type().Elem() {
		for i := 0; i < dataVal.Len(); i++ {
			val := dataVal.Index(i)
			channelVal.Send(val)
		}
		channelVal.Close()
	} else {
		Printfln("Unexpected types: %v, %v",
			channelVal.Type(), dataVal.Type())
	}
}

func main() {
	values := []string{"Alice", "Bob", "Charlie", "Dora"}
	channel := make(chan string)
	go sendOverChannel(channel, values)
	for {
		if val, open := <-channel; open {
			Printfln("Received value: %v", val)
		} else {
			break
		}
	}
}
