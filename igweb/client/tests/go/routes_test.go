package main

import (
	"strings"

	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	"github.com/gopherjs/gopherjs/js"
)

var wait = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Call("waitForSelector", "#primaryContent")
	return nil
})

func main() {

	routes := []string{"/", "/index", "/products", "/product-detail/swiss-army-knife", "/about", "/contact", "/shopping-cart"}
	tokenMap := map[string]string{"": "IGWEB", "/": "IGWEB", "/index": "IGWEB", "/products": "Add To Cart", "/product-detail/swiss-army-knife": "Swiss Army Knife", "/about": "Molly", "/contact": "Contact", "/shopping-cart": "Shopping Cart"}

	casper := js.Global.Get("casper")

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Routes Test", 7, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080", wait)
	})

	// Test each route on the website, and verify that the expected token string is found in the rendered content
	for _, r := range routes {
		func(r string) {
			casper.Call("then", func() {
				casper.Call("click", "a[href^='"+r+"']")
			})

			casper.Call("then", wait)

			casper.Call("then", func() {
				routeName := strings.Replace(r, `/`, "", -1)
				screenshotName := "route_render_test_" + routeName + ".png"
				casper.Call("capture", "screenshots/"+screenshotName)
				casper.Get("test").Call("assertTextExists", tokenMap[r], "Expected text \""+tokenMap[r]+"\", in body of web page, when accessing route: "+r)
			})
		}(r)
	}

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
