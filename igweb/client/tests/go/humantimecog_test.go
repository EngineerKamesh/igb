package main

import (
	"time"

	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	humanize "github.com/dustin/go-humanize"
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

	casper.Get("test").Call("begin", "Time Ago Cog Test Suite", 1, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/about", wait)
	})

	// Verify the human time representation of Molly's start date
	casper.Call("then", func() {
		mollysStartDate := time.Date(2017, 5, 24, 17, 9, 0, 0, time.UTC)
		mollysStartDateInHumanTime := humanize.Time(mollysStartDate)
		casper.Call("capture", "screenshots/timeago_cog_test.png")
		casper.Get("test").Call("assertSelectorHasText", "#Gopher-Molly .timeagoSpan", mollysStartDateInHumanTime, "Verify human time of Molly's start date produced by the Time Ago Cog.")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
