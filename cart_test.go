package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getCartSize_OneProduct(t *testing.T) {
	// Arrange
	var cartItems []CartItem
	cartItems = append(cartItems, CartItem{"0000001", 5})
	expectedSize := 5

	// Act
	cartSize := getCartSize(cartItems)

	// Assert
	require.Equal(t, expectedSize, cartSize)
}

func Test_getCartSize_TwoProducts(t *testing.T) {
	// Arrange
	var cartItems []CartItem
	cartItems = append(cartItems, CartItem{"0000001", 1})
	cartItems = append(cartItems, CartItem{"0000002", 3})
	expectedSize := 4

	// Act
	cartSize := getCartSize(cartItems)

	// Assert
	require.Equal(t, expectedSize, cartSize)
}
