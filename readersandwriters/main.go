package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(m)

	fmt.Print(writer.String())
}
