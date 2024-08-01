package main

import (
	"reflect"
	"strings"
)

func createMap(slice interface{}, op func(interface{}) interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() == reflect.Slice {
		mapType := reflect.MapOf(sliceVal.Type().Elem(), sliceVal.Type().Elem())
		mapVal := reflect.MakeMap(mapType)
		for i := 0; i < sliceVal.Len(); i++ {
			elemVal := sliceVal.Index(i)
			mapVal.SetMapIndex(elemVal, reflect.ValueOf(op(elemVal.Interface())))
		}
		return mapVal.Interface()
	}
	return nil
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	reverse := func(val interface{}) interface{} {
		if str, ok := val.(string); ok {
			return strings.ToUpper(str)
		}
		return val
	}
	namesMap := createMap(names, reverse).(map[string]string)
	for k, v := range namesMap {
		Printfln("Key: %v, Value:%v", k, v)
	}
}
