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

var casper = js.Global.Get("casper")

func main() {

	routesTokenMap := map[string]string{"/": "IGWEB", "/index": "IGWEB", "/products": "Add To Cart", "/product-detail/swiss-army-knife": "Swiss Army Knife", "/about": "Molly", "/contact": "Contact", "/shopping-cart": "Shopping Cart"}

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Client-Side Routes Test Suite", 7, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080", wait)
	})

	// Test each route on the website, and verify that the expected token string is found in the rendered content
	for route, expectedString := range routesTokenMap {
		func(route, expectedString string) {
			casper.Call("then", func() {
				casper.Call("click", "a[href^='"+route+"']")
			})

			casper.Call("then", func() {
				casper.Call("wait", 1800, func() {
					routeName := strings.Replace(route, `/`, "", -1)
					screenshotName := "route_render_test_" + routeName + ".png"
					casper.Call("capture", "screenshots/"+screenshotName)
					casper.Get("test").Call("assertTextExists", expectedString, "Expected text \""+expectedString+"\", in body of web page, when accessing route: "+route)
				})
			})
		}(route, expectedString)
	}

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
