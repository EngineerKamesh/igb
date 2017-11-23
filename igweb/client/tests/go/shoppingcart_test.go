package main

import (
	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	"github.com/gopherjs/gopherjs/js"
)

var wait = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Call("waitForSelector", "#primaryContent")
	return nil
})

var casper = js.Global.Get("casper")

func main() {

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Shopping Cart Test Suite", 3, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/products", wait)
	})

	// Add 3 Swiss Army Knives to the Shopping Cart
	for i := 0; i < 3; i++ {
		casper.Call("then", func() {
			casper.Call("click", ".addToCartButton:first-child")
		})
	}

	// Navigate to Shopping Cart page
	casper.Call("then", func() {
		casper.Call("click", "a[href^='/shopping-cart']")

	})

	// Verify product added and quantity
	casper.Call("wait", 207, func() {
		casper.Get("test").Call("assertTextExists", "Swiss Army Knife", "Display correct product in shopping cart.")
	})

	casper.Call("wait", 108, func() {
		casper.Get("test").Call("assertTextExists", "Quantity: 3", "Display correct product quantity in shopping cart.")
	})

	casper.Call("wait", 450, func() {
		casper.Call("capture", "screenshots/shoppingcart_test_add_item.png")
	})

	// Remove product from shopping cart
	casper.Call("then", func() {
		casper.Call("click", ".removeFromCartButton:first-child")
	})

	// After removing the product from the shopping cart, the shopping cart should be empty
	casper.Call("wait", 5004, func() {
		casper.Call("capture", "screenshots/shoppingcart_test_empty.png")
		casper.Get("test").Call("assertTextExists", "Your shopping cart is empty.", "Empty the shopping cart.")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
