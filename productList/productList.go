package productlist

import (
	"errors"
	"fmt"
	"math/rand"
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

func GetAdvertising(productName string) (string, error) {

	if productName == "" {
		return "", errors.New("Product name is empty!")
	} else if productName != "Bread" && productName != "Milk" && productName != "Sugar" {
		return "", errors.New("Product name does not exist!")
	}

	message := fmt.Sprintf(formatRandom(), productName)
	return message, nil

}

func formatRandom() string {
	advertise := []string{"%v is the best product",
		"%v is the tastiest product",
		"Today, we have discount %v for product"}
	return advertise[rand.Intn(len(advertise))]

}
