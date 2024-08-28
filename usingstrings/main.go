package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
func main() {
	name := "Kayak"
	Printfln("String: %s", name)
	Printfln("Character: %c", []rune(name)[0])
	Printfln("Unicode: %U", []rune(name)[0])
}
