package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"products": getProducts(),
	}); err != nil {
		panic(err.Error())
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	product, err := getProduct(productId)
	if err != nil {
		panic(err.Error())
	}
	if err := templates.ExecuteTemplate(w, "product", map[string]interface{}{
		"product": product,
	}); err != nil {
		panic(err.Error())
	}
}

func viewCartHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSessionId(r)
	cart, err := getCart(sessionId)
	if err != nil {
		panic(err.Error())
	}
	if err := templates.ExecuteTemplate(w, "cart", map[string]interface{}{
		"cart_size": cartSize(cart),
		//"shipping_cost": shippingCost,
		//"total_cost":    totalPrice,
		//"items":         items,
	}); err != nil {
		panic(err.Error())
	}
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSessionId(r)
	err := addToCart(sessionId, "1", 1)
	if err != nil {
		panic(err.Error())
	}
	cart, err := getCart(sessionId)
	if err != nil {
		panic(err.Error())
	}
	if err := templates.ExecuteTemplate(w, "cart", map[string]interface{}{
		"cart_size": cartSize(cart),
		//"shipping_cost": shippingCost,
		//"total_cost":    totalPrice,
		//"items":         items,
	}); err != nil {
		panic(err.Error())
	}
}

func emptyCartHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSessionId(r)
	err := emptyCart(sessionId)
	if err != nil {
		panic(err.Error())
	}
	cart, err := getCart(sessionId)
	if err != nil {
		panic(err.Error())
	}
	if err := templates.ExecuteTemplate(w, "cart", map[string]interface{}{
		"cart_size": cartSize(cart),
		//"shipping_cost": shippingCost,
		//"total_cost":    totalPrice,
		//"items":         items,
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

func getSessionId(r *http.Request) string {
	sessionId := r.Context().Value(keySessionID{})
	if sessionId != nil {
		return sessionId.(string)
	}
	return ""
}

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

func cartSize(c []CartItem) int {
	cartSize := 0
	for _, item := range c {
		cartSize += int(item.GetQuantity())
	}
	return cartSize
}
