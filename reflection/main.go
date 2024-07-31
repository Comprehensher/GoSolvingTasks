package main

func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
		case Product:
			Printfln("Product: Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		case Customer:
			Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
		}
	}
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer{Name: "Alice", City: "New York"}
	printDetails(product, customer)
}
