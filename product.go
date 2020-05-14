package shoppingCart

type Product struct {
	Title    string
	Price    float64
	Category *Category
}

func NewProduct(title string, price float64, category *Category) *Product {
	return &Product{Title: title, Price: price, Category: category}
}

func (p *Product) SetCategory(c *Category) {
	p.Category = c
}

func (p *Product) SetTitle(t string) {
	p.Title = t
}

func (p *Product) SetPrice(price float64) {
	p.Price = price
}
