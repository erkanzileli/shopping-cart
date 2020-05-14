package shoppingCart

type DeliveryCost struct {
	FixedCost       float64
	CostForDelivery float64
	CostForProduct  float64
}

func NewDeliveryCost(fixedCost float64, costForDelivery float64, costForProduct float64) *DeliveryCost {
	return &DeliveryCost{FixedCost: fixedCost, CostForDelivery: costForDelivery, CostForProduct: costForProduct}
}

func (d *DeliveryCost) CalculateFor(c *Cart) float64 {
	categories := make(map[*Category]bool)
	numberOfProducts := 0

	for _, v := range c.Items {
		_, ok := categories[v.Product.Category]
		if !ok {
			categories[v.Product.Category] = true
		}
		numberOfProducts++
	}

	return d.CostForDelivery*float64(len(categories)) + d.CostForProduct*float64(numberOfProducts) + d.FixedCost
}
