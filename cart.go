package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type Cart struct {
	ClientId  string     `json:"ClientId"`
	CartItems []CartItem `json:"CartItems"`
}

type CartItem struct {
	ProductId string `json:"ProductId"`
	Quantity  int    `json:"Quantity"`
}

func (cartItem *CartItem) GetQuantity() int {
	return cartItem.Quantity
}

func getCartItems(userId string) ([]CartItem, error) {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	var cart Cart
	_, err := client.
		R().
		SetResult(&cart).
		Get(addr + "/cart/" + userId)
	return cart.CartItems, err
}

func addToCart(userId string, productId string, quantity int) error {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	quantityStr := strconv.Itoa(quantity)
	var cart Cart
	_, err := client.
		R().
		SetFormData(map[string]string{
			"userId":    userId,
			"productId": productId,
			"quantity":  quantityStr,
		}).
		SetResult(&cart).
		Post(addr + "/cart")
	fmt.Printf("addToCartHandler %s %s", cart.ClientId, cart.CartItems[0].ProductId)
	return err
}

func emptyCart(userId string) error {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	_, err := client.
		R().
		Get(addr + "/cart/" + userId + "/empty")
	return err
}
