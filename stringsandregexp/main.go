package main

import (
	"fmt"
	"strings"
)

func main() {
	username := " Alice"
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed:", ">>"+trimmed+"<<")
}
