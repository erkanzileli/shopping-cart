package shoppingCart

type DiscountType int

const (
	Rate DiscountType = iota
	Amount
)

type Campaign struct {
	Quantity     int
	Amount       float64
	Category     *Category
	DiscountType DiscountType
}

func NewCampaign(category *Category, amount float64, quantity int, discountType DiscountType) *Campaign {
	return &Campaign{Category: category, Amount: amount, Quantity: quantity, DiscountType: discountType}
}
