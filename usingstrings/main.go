package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
func main() {
	name := "Kayak"
	Printfln("Bool: %t", len(name) > 1)
	Printfln("Bool: %t", len(name) > 100)
}
