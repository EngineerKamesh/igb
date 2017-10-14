package button

import (
	"reflect"

	"github.com/uxtoolkit/cog"
)

var cogType reflect.Type

type Button struct {
	cog.UXCog
}

func NewButton() *Button {

	myButton := &Button{}
	myButton.SetCogType(cogType)
	return myButton
}

func init() {
	cogType = reflect.TypeOf(Button{})
}
