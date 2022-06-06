package main

import (
	"testing"
)

func Test_GetNanos(t *testing.T) {
	price := Price{19, 990000000}
	expectedNanos := 990000000
	if expectedNanos != price.GetNanos() {
		t.Errorf("Price should be %d, got %d",
			expectedNanos, price.GetNanos())
	}
}

func Test_GetUnits(t *testing.T) {
	price := Price{19, 990000000}
	expectedUnits := 19
	if expectedUnits != price.GetUnits() {
		t.Errorf("Price should be %d, got %d",
			expectedUnits, price.GetUnits())
	}
}

func Test_GetFloat(t *testing.T) {
	price := Price{19, 990000000}
	expectedPrice := 19.99
	if expectedPrice != price.GetFloat() {
		t.Errorf("Price should be %f, got %f",
			expectedPrice, price.GetFloat())
	}
}

func Test_GetPrice(t *testing.T) {
	value := 19.99
	expectedPrice := Price{19, 990000000}
	price := GetPrice(value)
	if expectedPrice.GetUnits() != price.GetUnits() {
		t.Errorf("Price Units should be %d, got %d",
			expectedPrice.GetUnits(), price.GetUnits())
	}

	if expectedPrice.GetNanos() != price.GetNanos() {
		t.Errorf("Price Nanos should be %d, got %d",
			expectedPrice.GetNanos(), price.GetNanos())
	}
}

func Test_GetProductsPrice(t *testing.T) {
	var cartItems []CartItem
	cartItem := CartItem{"0000001", 1}
	cartItems = append(cartItems, cartItem)
	expectedPrice := Price{19, 990000000}
	price := GetProductsPrice(cartItems)

	if expectedPrice.GetUnits() != price.GetUnits() {
		t.Errorf("Price Units should be %d, got %d",
			expectedPrice.GetUnits(), price.GetUnits())
	}

	if expectedPrice.GetNanos() != price.GetNanos() {
		t.Errorf("Price Nanos should be %d, got %d",
			expectedPrice.GetNanos(), price.GetNanos())
	}
}

func Test_GetProductsPriceTwo(t *testing.T) {
	var cartItems []CartItem
	cartItem := CartItem{"0000001", 2}
	cartItems = append(cartItems, cartItem)
	expectedPrice := Price{39, 980000000}
	price := GetProductsPrice(cartItems)

	if expectedPrice.GetUnits() != price.GetUnits() {
		t.Errorf("Price Units should be %d, got %d",
			expectedPrice.GetUnits(), price.GetUnits())
	}

	if expectedPrice.GetNanos() != price.GetNanos() {
		t.Errorf("Price Nanos should be %d, got %d",
			expectedPrice.GetNanos(), price.GetNanos())
	}
}

func Test_GetTotalPrice(t *testing.T) {
	shippingPrice := Price{20, 000000000}
	productsPrice := Price{19, 990000000}
	expectedTotal := 39.99
	totalPrice := GetTotalPrice(shippingPrice, productsPrice)
	if expectedTotal != totalPrice.GetFloat() {
		t.Errorf("Price should be %f, got %f",
			expectedTotal, totalPrice.GetFloat())
	}
}
