package main

import (
	"math"
	"os"

	"github.com/go-resty/resty/v2"
)

// STRUCTS
type Products struct {
	Products []Product `json:"Products"`
}

type Product struct {
	ProductId   string   `json:"ProductId,omitempty"`
	Name        string   `json:"Name,omitempty"`
	Description string   `json:"Description"`
	PictureUrl  string   `json:"PictureUrl"`
	Price       Price    `json:"Price,omitempty"`
	Categories  []string `json:"Categories"`
}

// API
func getAllProducts() []Product {
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

func getProducts(productIds []string) ([]Product, error) {
	var products []Product
	for _, productId := range productIds {
		product, err := getProduct(productId)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func renderProductImage(productId string) string {
	product, _ := getProduct(productId)
	return product.PictureUrl
}

func renderProductName(productId string) string {
	product, _ := getProduct(productId)
	return product.Name
}

func renderProductPrice(productId string) string {
	product, _ := getProduct(productId)
	return renderPrice(product.Price)
}

// INNER FUNCTIONS
func getProductIds(cartItems []CartItem) []string {
	var productIds []string
	for i := range cartItems {
		productIds = append(productIds, cartItems[i].ProductId)
	}
	return productIds
}

func getPrice(fullValue float64) Price {
	units := int(math.Floor(fullValue))
	nanos := int(math.Round((fullValue-math.Floor(fullValue))*100) * 10000000)
	return Price{units, nanos}
}

func getProductsPrice(products []Product, cartItems []CartItem) Price {
	var totalValue float64
	for x := range products {
		for y := range cartItems {
			if products[x].ProductId == cartItems[y].ProductId {
				totalValue += products[x].Price.GetFloat() * float64(cartItems[y].Quantity)
			}
		}
	}
	return getPrice(totalValue)
}

func getTotalPrice(shippingPrice Price, productsPrice Price) Price {
	totalPrice := shippingPrice.GetFloat() + productsPrice.GetFloat()
	totalCost := getPrice(totalPrice)
	return totalCost
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
