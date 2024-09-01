package main

import (
	"io"
	"strings"
)

func main() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)
}
