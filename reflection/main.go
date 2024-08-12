package main

import (
	"reflect"
)

func createChannelAndSend(data interface{}) interface{} {
	dataVal := reflect.ValueOf(data)
	channelType := reflect.ChanOf(reflect.BothDir,
		dataVal.Type().Elem())
	channel := reflect.MakeChan(channelType, 1)
	go func() {
		for i := 0; i < dataVal.Len(); i++ {
			channel.Send(dataVal.Index(i))
		}
		channel.Close()
	}()
	return channel.Interface()
}

func main() {
	values := []string{"Alice", "Bob", "Charlie", "Dora"}
	channel := createChannelAndSend(values).(chan string)
	for {
		if val, open := <-channel; open {
			Printfln("Received value: %v", val)
		} else {
			break
		}
	}
}
