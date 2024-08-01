package main

import (
	"reflect"
)

func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		iter := mapValue.MapRange()
		for iter.Next() {
			Printfln("Map Key: %v, Value: %v", iter.Key(), iter.Value())
		}
	} else {
		Printfln("Not a map")
	}
}

func main() {
	pricesMap := map[string]float64{"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50}
	printMapContents(pricesMap)
}
