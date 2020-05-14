package shoppingCart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCampaignRate(t *testing.T) {
	cat := &Category{Title: "Food"}
	a := 10.11
	q := 5
	dt := Rate

	c := NewCampaign(cat, a, q, dt)

	assert.Equal(t, cat, c.Category)
	assert.Equal(t, a, c.Amount)
	assert.Equal(t, q, c.Quantity)
	assert.Equal(t, dt, c.DiscountType)
}

func TestNewCampaignAmount(t *testing.T) {
	cat := &Category{Title: "Food"}
	a := 10.11
	q := 5
	dt := Amount

	c := NewCampaign(cat, a, q, dt)

	assert.Equal(t, cat, c.Category)
	assert.Equal(t, a, c.Amount)
	assert.Equal(t, q, c.Quantity)
	assert.Equal(t, dt, c.DiscountType)
}
