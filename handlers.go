package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"products": getProducts(),
	}); err != nil {
		panic(err.Error())
	}
}

var (
	templates = template.Must(
		template.New("").Funcs(
			template.FuncMap{
				"renderPrice": renderPrice,
			}).ParseGlob("templates/*.html"))
)

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
