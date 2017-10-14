package carousel

import (
	"reflect"

	"github.com/gopherjs/gopherjs/js"
	"github.com/uxtoolkit/cog"
)

const (
	PLAYTEXT = "&#9658;"
	STOPTEXT = "&#9724;"
)

var cogType reflect.Type
var JS = js.Global

type CarouselParams struct {
	*js.Object
	Container    string   `js:"container"`
	Items        int      `js:"items"`
	SlideBy      string   `js:"slideBy"`
	Autoplay     bool     `js:"autoplay"`
	AutoplayText []string `js:"autoplayText"`
}

type Carousel struct {
	cog.UXCog
	carousel *js.Object
}

func NewCarousel() *Carousel {

	c := &Carousel{}
	c.SetCogType(cogType)
	return c
}

func (c *Carousel) Start() {

	if c.Props["contentItems"] == nil || c.Props["carouselContentID"] == nil {
		return
	}

	params := &CarouselParams{Object: js.Global.Get("Object").New()}
	params.Container = "#" + c.Props["carouselContentID"].(string)

	if c.Props["items"] == nil {
		params.Items = 1
	} else {
		params.Items = c.Props["items"].(int)
	}

	if c.Props["slideBy"] == nil {
		params.SlideBy = "page"
	} else {
		params.SlideBy = c.Props["slideBy"].(string)
	}

	if c.Props["autoplay"] == nil {
		params.Autoplay = true
	} else {
		params.Autoplay = c.Props["autoplay"].(bool)
	}

	if c.Props["autoplayText"] == nil {
		params.AutoplayText = []string{PLAYTEXT, STOPTEXT}
	} else {
		params.AutoplayText = c.Props["autoplayText"].([]string)
	}

	c.carousel = JS.Get("tns").New(params)

}

func init() {
	cogType = reflect.TypeOf(Carousel{})
}
