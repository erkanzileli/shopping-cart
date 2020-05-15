package shoppingCart

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCart(t *testing.T) {
	c := NewCart()
	assert.Equal(t, len(c.Items), 0)
}

func TestCart_AddItem(t *testing.T) {
	p1 := &Product{Title: "Chocolate"}
	p2 := &Product{Title: "Cocoa"}

	c := NewCart()

	err := c.AddItem(p1, 0)
	assert.Error(t, MinimumQuantityError, err)
	assert.Equal(t, 0, len(c.Items))

	err = c.AddItem(p1, -1)
	assert.Error(t, MinimumQuantityError, err)
	assert.Equal(t, 0, len(c.Items))

	err = c.AddItem(p1, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.Items))

	err = c.AddItem(p1, 2)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.Items))

	err = c.AddItem(p2, 0)
	assert.Error(t, MinimumQuantityError, err)
	assert.Equal(t, 1, len(c.Items))

	err = c.AddItem(p2, -1)
	assert.Error(t, MinimumQuantityError, err)
	assert.Equal(t, 1, len(c.Items))

	err = c.AddItem(p2, 1)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(c.Items))

	err = c.AddItem(p2, 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(c.Items))
}

func TestCart_GetItemCount(t *testing.T) {
	p1 := &Product{Title: "Chocolate"}
	c := NewCart()

	count, err := c.GetItemCount(p1)
	assert.Error(t, ProductNotExistError, err)
	assert.Equal(t, 0, count)

	_ = c.AddItem(p1, 2)

	count, err = c.GetItemCount(p1)
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

func TestCart_RemoveItem(t *testing.T) {
	cat := NewCategory("Food")
	p1 := NewProduct("Chocolate", 10.11, cat)
	p2 := NewProduct("Cocoa", 10.12, cat)

	c := NewCart()

	_ = c.AddItem(p1, 2)
	_ = c.AddItem(p2, 2)

	err := c.RemoveItem(p1, 0)
	assert.Error(t, MinimumQuantityError, err)

	err = c.RemoveItem(p1, -1)
	assert.Error(t, MinimumQuantityError, err)

	err = c.RemoveItem(p1, 3)
	assert.Error(t, ProductNotEnoughError, err)

	err = c.RemoveItem(p1, 1)
	assert.Nil(t, err)

	err = c.RemoveItem(p1, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.Items))

	err = c.RemoveItem(p1, 1)
	assert.Error(t, ProductNotExistError)

	err = c.RemoveItem(p2, 2)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(c.Items))

}

func TestCart_ApplyDiscounts(t *testing.T) {
	cat := NewCategory("Food")
	cat1 := NewCategory("Sport")
	p1 := NewProduct("Chocolate", 15, cat)
	p2 := NewProduct("Cocoa", 20, cat)
	cam1 := NewCampaign(cat, 10.0, 5, Rate)
	cam2 := NewCampaign(cat, 5.0, 3, Amount)
	cam3 := NewCampaign(cat1, 5.0, 3, Amount)

	c := NewCart()

	_ = c.AddItem(p1, 3)
	_ = c.AddItem(p2, 2)

	c.ApplyDiscounts(cam3)
	assert.Equal(t, 0.0, c.CampaignDiscount)

	c.ApplyDiscounts(cam1, cam2)
	want := (3*15 + 2*20) * 0.1

	assert.Equal(t, want, c.CampaignDiscount)
}

func TestCart_ApplyCoupon(t *testing.T) {
	cat := NewCategory("Food")
	p1 := NewProduct("Chocolate", 15, cat)
	p2 := NewProduct("Cocoa", 20, cat)
	cam1 := NewCampaign(cat, 5.1, 3, Amount)
	cp1 := NewCoupon(100.0, 10.0, Rate)
	cp2 := NewCoupon(100.0, 10.0, Amount)

	c := NewCart()

	_ = c.AddItem(p1, 3)
	_ = c.AddItem(p2, 3)

	c.ApplyCoupon(cp1)
	want := (p1.Price*3 + p2.Price*3) * cp1.Amount / 100
	assert.Equal(t, want, c.CouponDiscount)

	c.ApplyCoupon(cp2)
	want = cp2.Amount
	assert.Equal(t, want, c.CouponDiscount)

	c.ApplyDiscounts(cam1)

	c.ApplyCoupon(cp1)
	assert.Equal(t, 0.0, c.CouponDiscount)
}

func TestCart_GetCouponDiscount(t *testing.T) {
	cat := NewCategory("Food")
	p1 := NewProduct("Chocolate", 15, cat)
	p2 := NewProduct("Cocoa", 20, cat)
	cp1 := NewCoupon(100.0, 10.0, Amount)

	c := NewCart()

	_ = c.AddItem(p1, 5)
	_ = c.AddItem(p2, 3)

	c.ApplyCoupon(cp1)

	want := cp1.Amount
	got := c.GetCouponDiscount()

	assert.Equal(t, want, got)
}

func TestCart_GetCampaignDiscount(t *testing.T) {
	cat := NewCategory("Food")
	p1 := NewProduct("Chocolate", 15, cat)
	p2 := NewProduct("Cocoa", 20, cat)
	cam1 := NewCampaign(cat, 10.0, 5, Rate)
	cam2 := NewCampaign(cat, 5.0, 3, Amount)

	c := NewCart()
	_ = c.AddItem(p1, 3)
	_ = c.AddItem(p2, 2)

	c.ApplyDiscounts(cam1, cam2)

	want := (p1.Price*3 + p2.Price*2) * 0.1
	got := c.GetCampaignDiscount()

	assert.Equal(t, want, got)
}

func TestCart_GetTotalAmountAfterDiscounts(t *testing.T) {
	cat := NewCategory("Food")
	p1 := NewProduct("Chocolate", 25, cat)
	p2 := NewProduct("Cocoa", 20, cat)
	cam1 := NewCampaign(cat, 15.0, 5, Amount)
	cp1 := NewCoupon(100.0, 10.0, Amount)

	c := NewCart()
	_ = c.AddItem(p1, 3)
	_ = c.AddItem(p2, 2)

	c.ApplyDiscounts(cam1)
	c.ApplyCoupon(cp1)

	want := (p1.Price*3 + p2.Price*2) - cam1.Amount - cp1.Amount
	assert.Equal(t, want, c.GetTotalAmountAfterDiscounts())
}

func TestCart_Print(t *testing.T) {
	cat := NewCategory("Food")
	p1 := NewProduct("Chocolate", 25, cat)
	p2 := NewProduct("Cocoa", 20, cat)
	cam1 := NewCampaign(cat, 15.0, 5, Amount)
	cp1 := NewCoupon(100.0, 10.0, Amount)

	c := NewCart()
	_ = c.AddItem(p1, 3)
	_ = c.AddItem(p2, 2)

	c.ApplyDiscounts(cam1)
	c.ApplyCoupon(cp1)

	discount := c.GetCampaignDiscount() + c.GetCouponDiscount()
	dc := *NewDeliveryCost(2.99, 2.5, 2)

	want := fmt.Sprintf(`CategoryName, ProductName, Quantity, UnitPrice, TotalPrice, TotalDiscount
Food, Chocolate, 3, 25, 75, %v
Food, Cocoa, 2, 20, 40, %v
Total: %v, Delivery Cost: %v`, discount, discount, c.GetTotalAmountAfterDiscounts(), c.GetDeliveryCost(dc))

	buf := &bytes.Buffer{}

	err := c.Print(dc, buf)
	assert.Nil(t, err)
	assert.Equal(t, want, buf.String())
}

func TestCart_GetDeliveryCost(t *testing.T) {
	cat := NewCategory("Food")
	p1 := NewProduct("Chocolate", 25, cat)
	p2 := NewProduct("Cocoa", 20, cat)

	c := NewCart()
	_ = c.AddItem(p1, 3)
	_ = c.AddItem(p2, 2)

	dc := *NewDeliveryCost(2.99, 2.5, 2)

	want := dc.CalculateFor(c)

	assert.Equal(t, want, c.GetDeliveryCost(dc))
}
