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

	casper.Get("test").Call("begin", "Notify Cog Test Suite", 1, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/products", wait)
	})

	// Add an item to the shopping cart
	casper.Call("then", func() {
		casper.Call("click", ".addToCartButton:nth-child(1)")
	})

	// Verify that the notification has been displayed
	casper.Call("wait", 450, func() {
		casper.Get("test").Call("assertSelectorHasText", "#alertify-logs .alertify-log-success", "Item added to cart", "Display Notify Cog when item added to shopping cart.")
	})

	casper.Call("wait", 450, func() {
		casper.Call("capture", "screenshots/notify_cog_test.png")
	})

	// Navigate to Shopping Cart page
	casper.Call("then", func() {
		casper.Call("click", "a[href^='/shopping-cart']")

	})

	// Remove product from shopping cart
	casper.Call("wait", 450, func() {
		casper.Call("click", ".removeFromCartButton:first-child")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})
}
