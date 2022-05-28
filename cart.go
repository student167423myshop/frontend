package main

import (
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

func getCart(userId string) ([]CartItem, error) {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	var cart []CartItem
	_, err := client.R().
		SetResult(&cart).
		Get(addr + "/api/v1/cart/" + userId)
	return cart, err
}

func addToCart(userId string, productId string, quantity int) error {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	var cart []CartItem
	_, err := client.R().
		SetResult(&cart).
		Get(addr + "/api/v1/cart/" + userId + "/add/" + productId + "/" + strconv.Itoa(quantity))
	return err
}

func emptyCart(userId string) error {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	var cart []CartItem
	_, err := client.R().
		SetResult(&cart).
		Get(addr + "/api/v1/cart/" + userId + "/empty")
	return err
}
