package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"products": getAllProducts(),
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
	cartItems, err := getCartItems(getSessionId(w, r))
	if err != nil {
		panic(err.Error())
	}

	shippingCost := getShippingCost()
	productIds := getProductIds(cartItems)
	products, err := getProducts(productIds)
	if err != nil {
		panic(err.Error())
	}

	productsPrice := getProductsPrice(products, cartItems)
	totalCost := getTotalPrice(shippingCost, productsPrice)

	if err := templates.ExecuteTemplate(w, "cart", map[string]interface{}{
		"cart_size":     getCartSize(cartItems),
		"shipping_cost": shippingCost,
		"total_cost":    totalCost,
		"items":         cartItems,
	}); err != nil {
		panic(err.Error())
	}
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	productId := r.FormValue("productId")
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	err := addToCart(getSessionId(w, r), productId, quantity)
	if err != nil {
		panic(err.Error())
	}
	renderCart(w, r)
}

func emptyCartHandler(w http.ResponseWriter, r *http.Request) {
	err := emptyCart(getSessionId(w, r))
	if err != nil {
		panic(err.Error())
	}
	renderCart(w, r)
}

func newOrderHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	streetAddress := r.FormValue("street_address")
	zipCode := r.FormValue("zip_code")
	city := r.FormValue("city")
	address := Address{email, streetAddress, zipCode, city}
	userId := getSessionId(w, r)
	order, err := newOrder(userId, address)
	if err != nil {
		panic(err.Error())
	}
	url := "/zamowienie/" + order.OrderId
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func viewOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	order, err := getOrder(orderId)
	if err != nil {
		panic(err.Error())
	}
	if err := templates.ExecuteTemplate(w, "order", map[string]interface{}{
		"order": order,
	}); err != nil {
		panic(err.Error())
	}
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
