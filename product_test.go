package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getProductIds(t *testing.T) {
	// Arrange
	var expectedProductIds []string
	expectedProductIds = append(expectedProductIds, "0000001")
	expectedProductIds = append(expectedProductIds, "0000002")
	expectedProductIds = append(expectedProductIds, "0000003")
	var cartItems []CartItem
	cartItems = append(cartItems, CartItem{expectedProductIds[0], 1})
	cartItems = append(cartItems, CartItem{expectedProductIds[1], 2})
	cartItems = append(cartItems, CartItem{expectedProductIds[2], 4})

	// Act
	productIds := getProductIds(cartItems)

	// Assert
	require.Equal(t, expectedProductIds, productIds)
}

func Test_getPrice(t *testing.T) {
	// Arrange
	value := 19.99
	expectedPrice := Price{19, 990000000}

	// Act
	price := getPrice(value)

	// Assert
	require.Equal(t, expectedPrice.GetUnits(), price.GetUnits())
	require.Equal(t, expectedPrice.GetNanos(), price.GetNanos())
}

func getTestProduct(productId string, price Price) Product {
	var categories []string
	product := Product{productId, "", "", "", price, categories}
	return product
}

func Test_GetProductsPrice(t *testing.T) {
	// Arrange
	productPrice := Price{19, 990000000}
	var products []Product
	products = append(products, getTestProduct("0000001", productPrice))
	var cartItems []CartItem
	cartItems = append(cartItems, CartItem{"0000001", 1})
	expectedPrice := productPrice

	// Act
	price := getProductsPrice(products, cartItems)

	// Assert
	require.Equal(t, expectedPrice.GetUnits(), price.GetUnits())
	require.Equal(t, expectedPrice.GetNanos(), price.GetNanos())
}

func Test_GetProductsPriceTwo(t *testing.T) {
	// Arrange
	productPrice := Price{19, 990000000}
	var products []Product
	products = append(products, getTestProduct("0000001", productPrice))
	var cartItems []CartItem
	cartItems = append(cartItems, CartItem{"0000001", 2})
	expectedPrice := Price{39, 980000000}

	// Act
	price := getProductsPrice(products, cartItems)

	// Assert
	require.Equal(t, expectedPrice.GetUnits(), price.GetUnits())
	require.Equal(t, expectedPrice.GetNanos(), price.GetNanos())
}

func Test_GetTotalPrice(t *testing.T) {
	// Arrange
	shippingPrice := Price{20, 000000000}
	productsPrice := Price{19, 990000000}
	expectedTotal := 39.99

	// Act
	totalPrice := getTotalPrice(shippingPrice, productsPrice)

	// Assert
	require.Equal(t, expectedTotal, totalPrice.GetFloat())
}

func Test_GetUnits(t *testing.T) {
	// Arrange
	price := Price{19, 990000000}
	expectedUnits := 19

	// Act
	units := price.GetUnits()

	// Assert
	require.Equal(t, expectedUnits, units)
}

func Test_GetNanos(t *testing.T) {
	// Arrange
	price := Price{19, 990000000}
	expectedNanos := 990000000

	// Act
	nanos := price.GetNanos()

	// Assert
	require.Equal(t, expectedNanos, nanos)
}

func Test_GetFloat(t *testing.T) {
	// Arrange
	price := Price{19, 990000000}
	expectedPrice := 19.99

	// Act
	priceFloat := price.GetFloat()

	// Assert
	require.Equal(t, expectedPrice, priceFloat)
}
