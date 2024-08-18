package main

import (
	"fmt"

	_ "eu.corp/hello/packages/data"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(4532))
}
