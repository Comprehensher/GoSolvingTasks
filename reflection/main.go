package main

import (
	"reflect"
)

func selectValue(data interface{}, index int) (result interface{}) {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() == reflect.Slice {
		result = dataVal.Index(index).Interface()
	}
	return
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	val := selectValue(names, 1).(string)
	Printfln("Selected: %v", val)
}
