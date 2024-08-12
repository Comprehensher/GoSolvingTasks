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
	for i := 0; i < fieldVal.Type().NumMethod(); i++ {
		method := fieldVal.Type().Method(i)
		Printfln("Interface Method: %v, Exported: %v",
			method.Name, method.PkgPath == "")
	}
	Printfln("--------")
	if fieldVal.Kind() == reflect.Interface {
		Printfln("Underlying Type: %v",
			fieldVal.Elem().Type())
		for i := 0; i < fieldVal.Elem().Type().NumMethod(); i++ {
			method := fieldVal.Elem().Type().Method(i)
			Printfln("Underlying Method: %v", method.Name)
		}
	}
}

func main() {
	getUnderlying(Wrapper{NamedItem: &Product{}}, "NamedItem")
}
