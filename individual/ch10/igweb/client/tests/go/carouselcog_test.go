package main

import (
	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	"github.com/gopherjs/gopherjs/js"
)

var wait = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Call("waitForSelector", "#carousel")
	return nil
})

var casper = js.Global.Get("casper")

func main() {

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Carousel Cog Test Suite", 1, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/index", wait)
	})

	// Verify that the carousel cog has been loaded.
	casper.Call("wait", 1800, func() {
		casper.Get("test").Call("assertResourceExists", "watch.jpg", "Display carousel cog.")
	})

	casper.Call("then", func() {
		casper.Call("capture", "screenshots/carousel_cog_test.png")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
