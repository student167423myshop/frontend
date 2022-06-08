package main

import (
	"html/template"
	"net/http"
	"strconv"

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
	product, err := getProduct(vars["productId"])
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
	renderCart(w, r)
}

func renderCart(w http.ResponseWriter, r *http.Request) {
	cartItems, err := getCartItems(getSessionId(r))
	if err != nil {
		panic(err.Error())
	}

	shippingCost := GetShippingCost()
	totalCost := GetTotalPrice(shippingCost, GetProductsPrice(cartItems))

	if err := templates.ExecuteTemplate(w, "cart", map[string]interface{}{
		"cart_size":     cartSize(cartItems),
		"shipping_cost": shippingCost,
		"total_cost":    totalCost,
		"items":         cartItems,
	}); err != nil {
		panic(err.Error())
	}
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	var productId = r.FormValue("productId")
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	err := addToCart(getSessionId(r), productId, quantity)
	if err != nil {
		panic(err.Error())
	}
	renderCart(w, r)
}

func emptyCartHandler(w http.ResponseWriter, r *http.Request) {
	err := emptyCart(getSessionId(r))
	if err != nil {
		panic(err.Error())
	}
	renderCart(w, r)
}

var (
	templates = template.Must(
		template.New("").Funcs(
			template.FuncMap{
				"renderPrice":        renderPrice,
				"renderProductImage": renderProductImage,
				"renderProductName":  renderProductName,
				"renderProductPrice": renderProductPrice,
			}).ParseGlob("templates/*.html"))
)
