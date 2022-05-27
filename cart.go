package main

import (
	"os"

	"github.com/go-resty/resty/v2"
)

type CartItem struct {
	productId string `json:"productId"`
	quantity  int    `json:"quantity"`
}

func (cartItem *CartItem) GetQuantity() int {
	return cartItem.quantity
}

func getCart(userID string) ([]CartItem, error) {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:0000"
	}
	var cart []CartItem
	_, err := client.R().
		SetResult(&cart).
		Get(addr + "/api/v1/cart/" + userID)
	return cart, err
}
