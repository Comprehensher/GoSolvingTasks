package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
func main() {
	var name string
	var category string
	var price float64
	source := "Lifejacket Watersports 48.95"
	n, err := fmt.Sscan(source, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name,
			category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
