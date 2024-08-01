package main

import (
	"reflect"
)

func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() == reflect.Map {
		for _, keyVal := range mapValue.MapKeys() {
			reflectedVal := mapValue.MapIndex(keyVal)
			Printfln("Map Key: %v, Value: %v", keyVal, reflectedVal)
		}
	} else {
		Printfln("Not a map")
	}
}

func main() {
	pricesMap := map[string]float64{"Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50}
	printMapContents(pricesMap)
}
