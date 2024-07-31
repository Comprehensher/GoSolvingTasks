package main

import (
	"reflect"
	"strings"
)

func incrementOrUpper(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if elemValue.Kind() == reflect.Ptr {
			elemValue = elemValue.Elem()
		}
		if elemValue.CanSet() {
			switch elemValue.Kind() {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}
			Printfln("Modified Value: %v", elemValue)
		} else {
			Printfln("Cannot set %v: %v", elemValue.Kind(), elemValue)
		}
	}
}

func main() {
	name := "Alice"
	price := 279
	city := "London"
	incrementOrUpper(&name, &price, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}
