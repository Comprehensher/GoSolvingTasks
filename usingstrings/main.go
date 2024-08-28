package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
func main() {
	number := 279.00
	Printfln("Decimalless with exponent: %b", number)
	Printfln("Decimal with exponent: %e", number)
	Printfln("Decimal without exponent: %f", number)
	Printfln("Hexadecimal: %x, %X", number, number)
}
