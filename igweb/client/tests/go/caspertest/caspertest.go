package caspertest

import "github.com/gopherjs/gopherjs/js"

type ViewportParams struct {
	*js.Object
	Width  int `js:"width"`
	Height int `js:"height"`
}
