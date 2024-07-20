package productlist

import (
	"errors"
)

func GetProductPrice(productName string) (int, error) {
	productMap := make(map[string]int)
	productMap["Bread"] = 25
	productMap["Milk"] = 85
	productMap["Sugar"] = 39

	if productName == "" {
		return 0, errors.New("Product name is empty!")
	} else if productName != "Bread" && productName != "Milk" && productName != "Sugar" {
		return 0, errors.New("Product name does not exist!")
	}

	return productMap[productName], nil

}
