package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"
	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	// splitsAfter := strings.SplitAfter(description, " ")
	// for _, x := range splitsAfter {
	// fmt.Println("SplitAfter >>" + x + "<<")
	// }
}
