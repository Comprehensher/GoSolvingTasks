package main

//"math/rand"
//"time"

func main() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlicesByName(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}
