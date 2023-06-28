package application_test

import (
	"github.com/helikson/arquitetura-hexagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = "disabled"
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}