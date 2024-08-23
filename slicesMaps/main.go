package main

import (
	"fmt"
	"reflect"
)

func main() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
}
