package main

import (
	"os"

	"github.com/go-resty/resty/v2"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       Price      `json:"price,omitempty"`
	Categories  []Category `json:"categories"`
}

type Price struct {
	Units int `json:"units"`
	Nanos int `json:"nanos"`
}

type Category struct {
	Name string
}

func getProducts() []Product {
	client := resty.New()
	addr := os.Getenv("PRODUCT_CATALOG_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost"
	}
	var products Products
	_, err := client.R().
		SetResult(&products).
		Get(addr + ":3550/api/v1/products")
	if err != nil {
		panic(err.Error())
	}

	return products.Products
}
