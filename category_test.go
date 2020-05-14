package shoppingCart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCategory(t *testing.T) {
	title := "Food"

	c := NewCategory(title)

	assert.Equal(t, title, c.Title)
}

func TestCategory_SetTitle(t *testing.T) {
	t1 := "Food"
	t2 := "Sport"

	c := NewCategory(t1)

	c.SetTitle(t2)

	assert.Equal(t, t2, c.Title)
}

func TestCategory_SetParent(t *testing.T) {
	p := &Category{Title: "Food"}

	c := NewCategory("Food")

	c.SetParent(p)

	assert.Equal(t, p, c.Parent)
}
