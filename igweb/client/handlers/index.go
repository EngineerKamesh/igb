package handlers

import (
	"context"
	"time"

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
		InitializeIndexPage(env)
	})
}

func InitializeIndexPage(env *common.Env) {

	// Carousel Cog
	c := carousel.NewCarousel()
	c.CogInit(env.TemplateSet)
	c.SetID("carousel")
	contentItems := []string{"/static/images/products/watch.jpg", "/static/images/products/shirt.jpg", "/static/images/products/coffeemug.jpg"}
	c.SetProp("contentItems", contentItems)
	c.SetProp("carouselContentID", "gophersContent")
	err := c.Start()
	if err != nil {
		println("Encountered the following error when attempting to start the carousel cog: ", err)
	}

	// Localtime Live Clock Cog
	localZonename, localOffset := time.Now().In(time.Local).Zone()
	lc := liveclock.NewLiveClock()
	lc.CogInit(env.TemplateSet)
	lc.SetID("myLiveClock")
	lc.SetProp("timeLabel", "Local Time")
	lc.SetProp("timezoneName", localZonename)
	lc.SetProp("timezoneOffset", localOffset)
	err = lc.Start()
	if err != nil {
		println("Encountered the following error when attempting to start the local liveclock cog: ", err)
	}

	// Chennai Live Clock Cog
	chennai := liveclock.NewLiveClock()
	chennai.CogInit(env.TemplateSet)
	chennai.SetID("chennaiLiveClock")
	chennai.SetProp("timeLabel", "Chennai")
	chennai.SetProp("timezoneName", "IST")
	chennai.SetProp("timezoneOffset", int(+5.5*3600))
	err = chennai.Start()
	if err != nil {
		println("Encountered the following error when attempting to start the chennai liveclock cog: ", err)
	}

	// Singapore Live Clock Cog
	singapore := liveclock.NewLiveClock()
	singapore.CogInit(env.TemplateSet)
	singapore.SetID("singaporeLiveClock")
	singapore.SetProp("timeLabel", "Singapore")
	singapore.SetProp("timezoneName", "SST")
	singapore.SetProp("timezoneOffset", int(+8.0*3600))
	err = singapore.Start()
	if err != nil {
		println("Encountered the following error when attempting to start the singapore liveclock cog: ", err)
	}

	// Hawaii Live Clock Cog
	hawaii := liveclock.NewLiveClock()
	hawaii.CogInit(env.TemplateSet)
	hawaii.SetID("hawaiiLiveClock")
	hawaii.SetProp("timeLabel", "Hawaii")
	hawaii.SetProp("timezoneName", "HDT")
	hawaii.SetProp("timezoneOffset", int(-10.0*3600))
	err = hawaii.Start()
	if err != nil {
		println("Encountered the following error when attempting to start the hawaii liveclock cog: ", err)
	}

}
