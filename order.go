package main

import (
	"os"

	"github.com/go-resty/resty/v2"
)

// STRUCTS
type Order struct {
	OrderId   string     `json:"OrderId"`
	ClientId  string     `json:"ClientId"`
	CartItems []CartItem `json:"CartItems"`
	Address   Address    `json:"Address"`
	TotalPaid Price      `json:"TotalPaid"`
}

type Address struct {
	Email         string `json:"Email"`
	StreetAddress string `json:"StreetAddress"`
	ZipCode       string `json:"ZipCode"`
	City          string `json:"City"`
}

// API
func getOrder(orderId string) (Order, error) {
	client := resty.New()
	addr := os.Getenv("ORDER_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7071"
	}
	var order Order
	_, err := client.
		R().
		SetResult(&order).
		Get(addr + "/order/" + orderId)
	return order, err
}

func newOrder(userId string, address Address) (Order, error) {
	client := resty.New()
	addr := os.Getenv("ORDER_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7071"
	}
	var order Order
	_, err := client.
		R().
		SetFormData(map[string]string{
			"email":          address.Email,
			"street_address": address.StreetAddress,
			"zip_code":       address.ZipCode,
			"city":           address.City,
		}).
		SetResult(&order).
		Post(addr + "/order/" + userId)
	return order, err
}
