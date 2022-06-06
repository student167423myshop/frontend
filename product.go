package main

import (
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
