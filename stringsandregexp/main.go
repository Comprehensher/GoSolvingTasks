package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "This  is  double  spaced"
	splits := strings.Fields(description)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}
}
