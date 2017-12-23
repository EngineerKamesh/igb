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

	casper.Get("test").Call("begin", "Date Picker Cog Test Suite", 1, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/contact", wait)
	})

	// Verify that the date picker is activated upon clicking the date input field
	casper.Call("then", func() {
		casper.Call("click", "#byDateInput")
		casper.Call("capture", "screenshots/datepicker_cog_test.png")
		casper.Get("test").Call("assertVisible", ".pika-single", "Display Datepicker Cog.")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})
}
