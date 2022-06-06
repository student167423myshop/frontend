package main

import (
	"fmt"
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

func GetShippingCost() Price {
	return Price{20, 000000000}
}

func renderCart(w http.ResponseWriter, r *http.Request) {
	sessionId := getSessionId(r)

	cartItems, err := getCartItems(sessionId)
	if err != nil {
		panic(err.Error())
	}

	shippingCost := GetShippingCost()
	productsCost := GetProductsPrice(cartItems)
	totalCost := GetTotalPrice(shippingCost, productsCost)

	if err := templates.ExecuteTemplate(w, "cart", map[string]interface{}{
		"cart_size":     cartSize(cartItems),
		"shipping_cost": shippingCost,
		"total_cost":    totalCost,
		"items":         cartItems,
	}); err != nil {
		panic(err.Error())
	}
}

func viewCartHandler(w http.ResponseWriter, r *http.Request) {
	renderCart(w, r)
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("             addToCartHandler")
	var productId = r.FormValue("productId")
	var quantityStr = r.FormValue("quantity")
	quantity, _ := strconv.Atoi(quantityStr)

	sessionId := getSessionId(r)

	err := addToCart(sessionId, productId, quantity)
	if err != nil {
		panic(err.Error())
	}

	renderCart(w, r)
}

func emptyCartHandler(w http.ResponseWriter, r *http.Request) {

	sessionId := getSessionId(r)

	err := emptyCart(sessionId)
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

func getSessionId(r *http.Request) string {
	sessionId := r.Context().Value(keySessionID{})
	if sessionId != nil {
		return sessionId.(string)
	}
	return "0000001" // TODO: Zwracanie poprawnego sessionId
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

func renderProductImage(productId string) string {
	product, _ := getProduct(productId)
	return product.Picture
}

func renderProductName(productId string) string {
	product, _ := getProduct(productId)
	return product.Name
}

func renderProductPrice(productId string) string {
	product, _ := getProduct(productId)
	return renderPrice(product.Price)
}
