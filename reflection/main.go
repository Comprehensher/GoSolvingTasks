package main

import (
	"reflect"
)

type Wrapper struct {
	NamedItem
}

func getUnderlying(item Wrapper, fieldName string) {
	itemVal := reflect.ValueOf(item)
	fieldVal := itemVal.FieldByName(fieldName)
	Printfln("Field Type: %v", fieldVal.Type())
	if fieldVal.Kind() == reflect.Interface {
		Printfln("Underlying Type: %v",
			fieldVal.Elem().Type())
	}
}

func main() {
	getUnderlying(Wrapper{NamedItem: &Product{}}, "NamedItem")
}
