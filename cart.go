package shoppingCart

import (
	"errors"
	"fmt"
)

var (
	ProductNotExistError  = errors.New("product doesn't exist")
	MinimumQuantityError  = errors.New("quantity can not be lower than 1")
	ProductNotEnoughError = errors.New("there aren't enough product")
)

type cartItem struct {
	Product  *Product
	Quantity int
}

type Cart struct {
	Items            []*cartItem
	CampaignDiscount float64
	CouponDiscount   float64
}

func NewCart() *Cart {
	return &Cart{
		Items: make([]*cartItem, 0),
	}
}

func (c *Cart) AddItem(p *Product, q int) error {
	if q < 1 {
		return MinimumQuantityError
	}
	for _, v := range c.Items {
		if v.Product == p {
			v.Quantity += q
			return nil
		}
	}
	c.Items = append(c.Items, &cartItem{p, q})
	return nil
}

func (c *Cart) RemoveItem(p *Product, q int) error {
	if q < 1 {
		return MinimumQuantityError
	}
	for i, v := range c.Items {
		if v.Product == p {
			if v.Quantity > q {
				v.Quantity -= q
			} else if v.Quantity == q {
				c.Items = append(c.Items[:i], c.Items[i+1:]...)
			} else {
				return ProductNotEnoughError
			}
		}
	}
	return nil
}

func (c *Cart) GetItemCount(p *Product) (int, error) {
	for _, v := range c.Items {
		if v.Product == p {
			return v.Quantity, nil
		}
	}
	return 0, ProductNotExistError
}

func (c *Cart) ApplyDiscounts(campaigns ...*Campaign) float64 {
	bigger := 0.0
	for _, v := range campaigns {
		discount := calculateCampaignDiscount(c.Items, v)
		if discount > bigger {
			bigger = discount
		}
	}
	c.CampaignDiscount = bigger
	return bigger
}

func (c *Cart) ApplyCoupon(coupon *Coupon) {
	totalAmount := 0.0
	discount := 0.0

	for _, v := range c.Items {
		totalAmount += v.Product.Price * float64(v.Quantity)
	}
	totalAmount -= c.CampaignDiscount

	if totalAmount >= coupon.MinAmount {
		if coupon.DiscountType == Rate {
			discount = totalAmount * coupon.Amount / 100
		} else {
			discount = coupon.Amount
		}
	}
	c.CouponDiscount = discount
}

// comes from cart responsibilities but we don't have any delivery parameters(costPerDelivery,costPerProduct) at this point
// so we can use cart.GetDeliveryCost function with this way :(
func (c *Cart) GetDeliveryCost(dc DeliveryCost) float64 {
	return dc.CalculateFor(c)
}

func (c *Cart) GetCampaignDiscount() float64 { return c.CampaignDiscount }

func (c *Cart) GetCouponDiscount() float64 { return c.CouponDiscount }

func (c *Cart) GetTotalAmountAfterDiscounts() float64 {
	totalAmount := 0.0
	for _, v := range c.Items {
		totalAmount += v.Product.Price * float64(v.Quantity)
	}
	return totalAmount - c.CampaignDiscount - c.CouponDiscount
}

func (c *Cart) Print(dc DeliveryCost) string {
	result := `CategoryName, ProductName, Quantity, UnitPrice, TotalPrice, TotalDiscount`
	discount := c.GetCampaignDiscount() + c.GetCouponDiscount()

	for _, v := range c.Items {
		result += fmt.Sprintf("\n%s, %s, %v, %v, %v, %v",
			v.Product.Category.Title, v.Product.Title, v.Quantity,
			v.Product.Price, float64(v.Quantity)*v.Product.Price, discount)
	}
	result += fmt.Sprintf("\nTotal: %v, Delivery Cost: %v", c.GetTotalAmountAfterDiscounts(), c.GetDeliveryCost(dc))

	return result
}

func calculateCampaignDiscount(items []*cartItem, cam *Campaign) float64 {
	amount := 0.0
	quantity := 0
	for _, v := range items {
		if v.Product.Category == cam.Category {
			amount += v.Product.Price * float64(v.Quantity)
			quantity += v.Quantity
		}
	}
	if quantity >= cam.Quantity {
		if cam.DiscountType == Rate {
			return amount * cam.Amount / 100
		} else {
			return cam.Amount
		}
	}
	return 0
}
