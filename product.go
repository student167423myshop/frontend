package main

import (
	"math"
	"os"

	"github.com/go-resty/resty/v2"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       Price    `json:"price,omitempty"`
	Categories  []string `json:"categories"`
}

type Price struct {
	Units int `json:"units"`
	Nanos int `json:"nanos"`
}

func (price *Price) GetUnits() int {
	if price != nil {
		return price.Units
	}
	return 0
}

func (price *Price) GetNanos() int {
	if price != nil {
		return price.Nanos
	}
	return 0
}

func (price *Price) GetFloat() float64 {
	units := float64(price.GetUnits())
	nanos := float64(price.GetNanos()) / 1000000000
	fullPrice := units + nanos
	return fullPrice
}

func getProducts() []Product {
	client := resty.New()
	addr := os.Getenv("PRODUCT_CATALOG_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:3550"
	}
	var products Products
	_, err := client.R().
		SetResult(&products).
		Get(addr + "/products")
	if err != nil {
		return nil
	}

	return products.Products
}

func getProduct(productId string) (Product, error) {
	client := resty.New()
	addr := os.Getenv("PRODUCT_CATALOG_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:3550"
	}
	var product Product
	_, err := client.R().
		SetResult(&product).
		Get(addr + "/product/" + productId)
	if err != nil {
		panic(err.Error())
	}

	return product, nil
}

func GetPrice(fullValue float64) Price {
	units := int(math.Floor(fullValue))
	nanos := int(math.Round((fullValue-math.Floor(fullValue))*100) * 10000000)
	return Price{units, nanos}
}

func GetProductsPrice(cartItems []CartItem) Price {
	var totalValue float64
	for _, cartItem := range cartItems {
		product, _ := getProduct(cartItem.ProductId)
		totalValue += product.Price.GetFloat() * float64(cartItem.Quantity)
	}
	return GetPrice(totalValue)
}

func GetTotalPrice(shippingPrice Price, productsPrice Price) Price {
	totalPrice := shippingPrice.GetFloat() + productsPrice.GetFloat()
	totalCost := GetPrice(totalPrice)
	return totalCost
}
