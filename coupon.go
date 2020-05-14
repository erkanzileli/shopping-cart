package shoppingCart

type Coupon struct {
	MinAmount    float64
	Amount       float64
	DiscountType DiscountType
}

func NewCoupon(minAmount, amount float64, discountType DiscountType) *Coupon {
	return &Coupon{MinAmount: minAmount, Amount: amount, DiscountType: discountType}
}

