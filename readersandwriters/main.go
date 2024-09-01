package main

import (
	"fmt"
	"io"
	"strings"
)

func scanFromReader(reader io.Reader, template string,
	vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func main() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(reader, scanTemplate, &name,
		&category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
}
