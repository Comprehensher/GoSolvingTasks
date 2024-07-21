package productlist

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"time"
)

type Product struct {
	Name         string
	Price        int
	ArrivingDate time.Time
}

var products = []Product{
	Product{
		Name:         "Bread",
		Price:        25,
		ArrivingDate: time.Date(2024, 8, 15, 14, 30, 45, 100, time.Local),
	},
	Product{
		Name:         "Bread",
		Price:        35,
		ArrivingDate: time.Date(2024, 8, 15, 14, 30, 45, 100, time.Local),
	},
	Product{
		Name:         "Milk",
		Price:        85,
		ArrivingDate: time.Date(2024, 8, 17, 14, 30, 45, 100, time.Local),
	},
	Product{
		Name:         "Sugar",
		Price:        39,
		ArrivingDate: time.Date(2024, 8, 18, 14, 30, 45, 100, time.Local),
	},
}

func (p Product) IsEmpty() bool {
	return p.Name == ""
}

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

func GetProductInfo(productName string) (Product, error) {
	for _, product := range products {
		if product.Name == productName {
			return product, nil
		}
	}
	return Product{}, errors.New("No found such product")
}

// Function with defer will be calles after outer method finishes or when exception
func GetProductWithPanic(productName string, price int) (Product, error) {
	if price <= 0 {
		panic("incorrect price")
	}

	for _, product := range products {
		if product.Name == productName && product.Price == price {
			return product, nil
		}
	}

	return Product{}, errors.New("no found such product")

}
