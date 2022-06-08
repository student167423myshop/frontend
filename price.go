package main

import "fmt"

func renderPrice(price Price) string {
	currencyLogo := renderCurrency()
	return fmt.Sprintf("%d.%02d %s",
		price.GetUnits(),
		price.GetNanos()/10000000,
		currencyLogo)
}

func renderCurrency() string {
	return "PLN"
}

func GetShippingCost() Price {
	return Price{20, 000000000}
}
