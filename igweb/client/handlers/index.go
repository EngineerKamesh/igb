package handlers

import (
	"context"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/carousel"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/liveclock"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/isomorphicgo/isokit"
)

func IndexHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		templateData := templatedata.Index{PageTitle: "IGWEB"}
		env.TemplateSet.Render("index_content", &isokit.RenderParams{Data: templateData, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: templateData.PageTitle})
		InitializeIndexEventHandlers(env)
	})
}

func InitializeIndexEventHandlers(env *common.Env) {

	// Carousel Cog
	c := carousel.NewCarousel()
	c.CogInit(env.TemplateSet)
	c.SetID("carousel")
	contentItems := []string{"/static/images/products/watch.jpg", "/static/images/products/shirt.jpg", "/static/images/products/coffeemug.jpg"}
	c.SetProp("contentItems", contentItems)
	c.SetProp("carouselContentID", "gophersContent")
	c.Render()
	c.Start()

	// Live Clock Cogs
	lc := liveclock.NewLiveClock()
	lc.CogInit(env.TemplateSet)
	lc.SetID("myLiveClock")
	lc.SetProp("timeLabel", "Local Time")
	lc.Render()
	lc.Start()

	chennai := liveclock.NewLiveClock()
	chennai.CogInit(env.TemplateSet)
	chennai.SetID("chennaiLiveClock")
	chennai.SetProp("timeLabel", "Chennai")
	chennai.SetProp("timezoneName", "IST")
	chennai.SetProp("timezoneOffset", int(+5.5*3600))
	chennai.Render()
	chennai.Start()

	singapore := liveclock.NewLiveClock()
	singapore.CogInit(env.TemplateSet)
	singapore.SetID("singaporeLiveClock")
	singapore.SetProp("timeLabel", "Singapore")
	singapore.SetProp("timezoneName", "SST")
	singapore.SetProp("timezoneOffset", int(+8.0*3600))
	singapore.Render()
	singapore.Start()

	hawaii := liveclock.NewLiveClock()
	hawaii.CogInit(env.TemplateSet)
	hawaii.SetID("hawaiiLiveClock")
	hawaii.SetProp("timeLabel", "Hawaii")
	hawaii.SetProp("timezoneName", "HDT")
	hawaii.SetProp("timezoneOffset", int(-9.0*3600))
	hawaii.Render()
	hawaii.Start()

}
