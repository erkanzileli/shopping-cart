package shoppingCart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCouponRate(t *testing.T) {
	min := 5.1
	a := 10.11
	dt := Rate

	c := NewCoupon(min, a, dt)

	assert.Equal(t, min, c.MinAmount)
	assert.Equal(t, a, c.Amount)
	assert.Equal(t, dt, c.DiscountType)
}

func TestNewCouponAmount(t *testing.T) {
	min := 5.1
	a := 10.11
	dt := Amount

	c := NewCoupon(min, a, dt)

	assert.Equal(t, min, c.MinAmount)
	assert.Equal(t, a, c.Amount)
	assert.Equal(t, dt, c.DiscountType)
}
