package main

import (
	"reflect"
)

func contains(slice interface{}, target interface{}) (found bool) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		for i := 0; i < sliceVal.Len(); i++ {
			if reflect.DeepEqual(sliceVal.Index(i).Interface(), target) {
				found = true
			}
		}
	}
	return
}

func main() {
	// name := "Alice"
	// price := 279
	city := "London"
	citiesSlice := []string{"Paris", "Rome", "London"}
	Printfln("Found #1: %v", contains(citiesSlice, city))
	sliceOfSlices := [][]string{citiesSlice, {"First", "Second", "Third"}}
	Printfln("Found #2: %v", contains(sliceOfSlices, citiesSlice))
}
