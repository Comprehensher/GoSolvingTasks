package main

import (
	"reflect"
)

func checkImplementation(check interface{}, targets ...interface{}) {
	checkType := reflect.TypeOf(check)
	if checkType.Kind() == reflect.Ptr &&
		checkType.Elem().Kind() == reflect.Interface {
		checkType := checkType.Elem()
		for _, target := range targets {
			targetType := reflect.TypeOf(target)
			Printfln("Type %v implements %v: %v",
				targetType, checkType,
				targetType.Implements(checkType))
		}
	}
}

func main() {
	currencyItemType := (*CurrencyItem)(nil)
	checkImplementation(currencyItemType, Product{}, &Product{}, &Purchase{})
}
