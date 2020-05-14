package shoppingCart

type Category struct {
	Title    string
	Parent   *Category
}

func NewCategory(title string) *Category {
	return &Category{Title: title}
}

func (c *Category) SetTitle(t string) {
	c.Title = t
}

func (c *Category) SetParent(parent *Category) {
	c.Parent = parent
}
