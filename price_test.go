package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_renderPrice(t *testing.T) {
	// Arrange
	price := Price{19, 990000000}
	expectedRender := "19.99 PLN"

	// Act
	renderedPrice := renderPrice(price)

	// Assert
	require.Equal(t, expectedRender, renderedPrice)
}

func Test_getShippingCost(t *testing.T) {
	// Arrange
	expectedPrice := Price{20, 000000000}

	// Act
	shippingPrice := getShippingCost()

	// Assert
	require.Equal(t, expectedPrice, shippingPrice)
}
