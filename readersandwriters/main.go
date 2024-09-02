package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	encoder.Encode(&dp)

	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)
	fmt.Print(writer.String())
}
