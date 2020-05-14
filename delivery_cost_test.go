package shoppingCart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeliveryCost(t *testing.T) {
	fixedCost := 2.99
	costForDelivery := 2.5
	costForProduct := 2.0
	dc := NewDeliveryCost(fixedCost, costForDelivery, costForProduct)

	assert.Equal(t, fixedCost, dc.FixedCost)
	assert.Equal(t, costForDelivery, dc.CostForDelivery)
	assert.Equal(t, costForProduct, dc.CostForProduct)
}

func TestDeliveryCost_CalculateFor(t *testing.T) {
	cat1 := NewCategory("Food")
	cat2 := NewCategory("Food2")
	p1 := NewProduct("Chocolate", 15, cat1)
	p2 := NewProduct("Cocoa", 20, cat2)
	p3 := NewProduct("Mini Cocoa", 15, cat2)

	c := NewCart()
	_ = c.AddItem(p1, 2)
	_ = c.AddItem(p2, 3)
	_ = c.AddItem(p3, 4)

	fixedCost := 2.99
	costForDelivery := 2.5
	costForProduct := 2.0
	dc := NewDeliveryCost(fixedCost, costForDelivery, costForProduct)

	want := costForDelivery*2 + costForProduct*3 + fixedCost

	got := dc.CalculateFor(c)

	assert.Equal(t, want, got)
}
