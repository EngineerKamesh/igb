package carousel

import (
	"errors"
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
	Controls     bool     `js:"controls"`
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

func (c *Carousel) Start() error {

	if c.Props["contentItems"] == nil || c.Props["carouselContentID"] == nil {
		return errors.New("The contentItems and carouselContentID props need to be set!")
	}

	err := c.Render()
	if err != nil {
		return err
	}

	params := &CarouselParams{Object: js.Global.Get("Object").New()}

	// Set the default parameter values
	params.Items = 1
	params.SlideBy = "page"
	params.Autoplay = true
	params.AutoplayText = []string{PLAYTEXT, STOPTEXT}
	params.Controls = false

	for propName, propValue := range c.Props {
		switch propName {

		case "carouselContentID":
			if propValue != nil {
				params.Container = "#" + c.Props["carouselContentID"].(string)
			}

		case "contentItems":
			// Do nothing

		case "items":
			if propValue != nil {
				params.Items = propValue.(int)
			}

		case "slideBy":
			if propValue != nil {
				params.SlideBy = c.Props["slideBy"].(string)
			}

		case "autoplay":
			if propValue != nil {
				params.Autoplay = c.Props["autoplay"].(bool)
			}

		case "autoplayText":
			if propValue != nil {
				params.AutoplayText = c.Props["autoplayText"].([]string)
			}

		case "controls":
			if propValue != nil {
				params.Controls = c.Props["controls"].(bool)
			}

		default:
			println("Warning: Unknown prop name provided: ", propName)
		}
	}
	c.carousel = JS.Get("tns").New(params)

	return nil
}

func init() {
	cogType = reflect.TypeOf(Carousel{})
}
