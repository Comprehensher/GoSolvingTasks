package main

import (
	"reflect"
)

func executeFirstVoidMethod(s interface{}) {
	sVal := reflect.ValueOf(s)
	for i := 0; i < sVal.NumMethod(); i++ {
		method := sVal.Type().Method(i)
		if method.Type.NumIn() == 1 {
			results := method.Func.Call([]reflect.Value{sVal})
			Printfln("Type: %v, Method: %v, Results: %v",
				sVal.Type(), method.Name, results)
			break
		} else {
			Printfln("Skipping method %v %v", method.Name,
				method.Type.NumIn())
		}
	}
}

func main() {
	executeFirstVoidMethod(&Product{Name: "Kayak", Price: 279})
}
