package main

import (
	"time"

	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	"github.com/gopherjs/gopherjs/js"
)

var wait = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Call("waitForSelector", "#myLiveClock div")
	return nil
})

var casper = js.Global.Get("casper")

func main() {

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Live Clock Cog Test Suite", 1, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/index", wait)
	})

	// Verify that the live clock shows the current time for the local time zone
	casper.Call("then", func() {
		casper.Call("wait", 900, func() {

			localZonename, localOffset := time.Now().In(time.Local).Zone()
			const layout = time.RFC1123
			var location *time.Location
			location = time.FixedZone(localZonename, localOffset)
			casper.Call("wait", 10, func() {
				t := time.Now()
				currentTime := t.In(location).Format(layout)
				casper.Get("test").Call("assertSelectorHasText", "#myLiveClock div", currentTime, "Display live clock for local timezone.")
			})

		})
	})

	casper.Call("then", func() {
		casper.Call("capture", "screenshots/liveclock_cog_test.png")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
