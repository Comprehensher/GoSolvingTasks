package productlist

func GetProductPrice(productName string) int {
	productMap := make(map[string]int)
	productMap["Bread"] = 25
	productMap["Milk"] = 85
	productMap["Sugar"] = 39

	return productMap[productName]

}
