package main

import (
	"fmt"
	"reflect"
)

func Printfln(msg string, vars ...interface{}) {
	fmt.Printf(msg+"\n", vars...)
}

type Wrapper struct {
	CatDetail
}

func executeMethod(src interface{}, reqMethod string, args ...interface{}) {
	srcT := reflect.TypeOf(src)
	structV := reflect.ValueOf(src)

	if srcT.Kind() == reflect.Struct || (srcT.Kind() == reflect.Pointer && srcT.Elem().Kind() == reflect.Struct) {
		for i := 0; i < srcT.NumMethod(); i++ {
			method := srcT.Method(i)
			if method.Name == reqMethod && method.Type.NumIn() == len(args)+1 {
				params := []reflect.Value{structV}
				match := true

				for j, arg := range args {
					param := reflect.ValueOf(arg)
					expectedType := method.Type.In(j + 1) // +1 to skip receiver
					if param.Type() != expectedType {
						Printfln("Argument %d has different type: expected %v, got %v", j, expectedType, param.Type())
						match = false
						break
					}
					params = append(params, param)
				}

				if match {
					outs := method.Func.Call(params)
					if len(outs) > 0 {
						Printfln("Res - %v", outs[0])
					}
					return
				}
			}
		}
		Printfln("No matching method found")
	}
}

func executeMethodVal(src interface{}, reqMethod string, args ...interface{}) {
	srcT := reflect.TypeOf(src)
	structV := reflect.ValueOf(src)

	if srcT.Kind() == reflect.Struct || (srcT.Kind() == reflect.Pointer && srcT.Elem().Kind() == reflect.Struct) {
		for i := 0; i < structV.NumMethod(); i++ {
			method := structV.Method(i)
			if srcT.Method(i).Name == reqMethod && method.Type().NumIn() == len(args) {
				params := []reflect.Value{}
				match := true

				for j, arg := range args {
					param := reflect.ValueOf(arg)
					expectedType := method.Type().In(j)
					if param.Type() != expectedType {
						Printfln("Argument %d has different type: expected %v, got %v", j, expectedType, param.Type())
						match = false
						break
					}
					params = append(params, param)
				}

				if match {
					outs := method.Call(params)
					if len(outs) > 0 {
						Printfln("Res - %v", outs[0])
					}
					return
				}
			}
		}
		Printfln("No matching method found")
	}
}

func main() {
	Sofa := Cat{Name: "Sofa", Height: 30}
	executeMethod(Sofa, "GrowUp", 4.02)
	executeMethodVal(&Sofa, "PortionToEat", 2.00, 4.00)
}
