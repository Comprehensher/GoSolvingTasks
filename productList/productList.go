package productlist

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
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

func GetProductsPrice(productNames []string) (map[string]int, error) {
	productMap := make(map[string]int)
	productMap["Bread"] = 25
	productMap["Milk"] = 85
	productMap["Sugar"] = 39

	resMap := make(map[string]int)

	if len(productNames) == 0 {
		return nil, errors.New("Product names are empty!")
	} else if !slices.Contains(productNames, "Bread") && !slices.Contains(productNames, "Milk") && !slices.Contains(productNames, "Sugar") {
		return nil, errors.New("Product name does not exist!")
	}

	for _, productName := range productNames {
		switch productName {
		case "Bread":
			resMap["Bread"] = productMap["Bread"]
		case "Milk":
			resMap["Milk"] = productMap["Milk"]
		case "Sugar":
			resMap["Sugar"] = productMap["Sugar"]
		}
	}

	return resMap, nil

}
