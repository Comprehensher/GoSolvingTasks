package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
func main() {
	name := "Kayak"
	Printfln("Pointer: %p", &name)
}
