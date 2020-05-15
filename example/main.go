package main

import (
	"log"
	"os"
	"shoppingCart"
)

func main() {
	cat1 := shoppingCart.NewCategory("Food")
	cat2 := shoppingCart.NewCategory("Sport")

	p1 := shoppingCart.NewProduct("Chocolate", 10, cat1)
	p2 := shoppingCart.NewProduct("Cocoa", 15, cat1)
	p3 := shoppingCart.NewProduct("Gloves", 45, cat2)
	p4 := shoppingCart.NewProduct("Dumbell", 75, cat2)

	camp1 := shoppingCart.NewCampaign(cat1, 5, 3, shoppingCart.Amount)
	camp2 := shoppingCart.NewCampaign(cat2, 5, 3, shoppingCart.Rate)

	cp := shoppingCart.NewCoupon(150, 10, shoppingCart.Rate)

	dc := *shoppingCart.NewDeliveryCost(2.99, 2, 1.5)

	c := shoppingCart.NewCart()

	err := c.AddItem(p1, 0)
	checkErr(err)

	err = c.AddItem(p1, -2)
	checkErr(err)

	err = c.AddItem(p1, 3)
	checkErr(err)

	_ = c.AddItem(p2, 2)
	_ = c.AddItem(p3, 1)
	_ = c.AddItem(p4, 1)

	c.ApplyDiscounts(camp1, camp2)
	c.ApplyCoupon(cp)

	err = c.Print(dc, os.Stdout)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
