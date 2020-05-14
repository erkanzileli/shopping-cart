package shoppingCart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	c := &Category{Title: "Food"}
	title := "Chocolate"
	price := 10.11

	p := NewProduct(title, price, c)

	assert.Equal(t, c, p.Category)
	assert.Equal(t, title, p.Title)
	assert.Equal(t, price, p.Price)
}

func TestProduct_SetCategory(t *testing.T) {
	c1 := &Category{Title: "Food"}
	c2 := &Category{Title: "Sport"}

	p := NewProduct("Chocolate", 10.11, c1)

	p.SetCategory(c2)

	assert.Equal(t, c2, p.Category)
}

func TestProduct_SetTitle(t *testing.T) {
	t1 := "Chocolate"
	t2 := "Cocoa"

	p := NewProduct(t1, 10.11, &Category{Title: "Food"})

	p.SetTitle(t2)

	assert.Equal(t, t2, p.Title)
}

func TestProduct_SetPrice(t *testing.T) {
	price1 := 10.11
	price2 := 10.12

	p := NewProduct("Chocolate", price1, &Category{Title: "Food"})

	p.SetPrice(price2)

	assert.Equal(t, price2, p.Price)
}
