package productlist

import (
	"testing"
)

func TestGetProductPrice(t *testing.T) {
	productName := "Milk"
	want := 85
	ER, error := GetProductPrice(productName)
	if ER != want || error != nil {
		t.Fatalf("Hello product %v have no price %v", productName, want)
	}
}

func TestProductEmptyPriceError(t *testing.T) {
	productName := ""
	price, error := GetProductPrice(productName)
	if error == nil {
		t.Fatalf("Hello, product has price %v want have error %v", price, error)
	}
}
