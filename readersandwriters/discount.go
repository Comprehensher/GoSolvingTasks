package main

type DiscountedProduct struct {
	*Product `json:"product,omitempty"`
	Discount float64 `json:"-"`
}
