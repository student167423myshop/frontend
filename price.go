package main

import "fmt"

// Structs
type Price struct {
	Units int `json:"Units"`
	Nanos int `json:"Nanos"`
}

// INNER FUNCTIONS
func renderPrice(price Price) string {
	currencyLogo := "PLN"
	return fmt.Sprintf("%d.%02d %s",
		price.GetUnits(),
		price.GetNanos()/10000000,
		currencyLogo)
}

func getShippingCost() Price {
	return Price{20, 000000000}
}
