package timeago

import (
	"errors"
	"reflect"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/uxtoolkit/cog"
)

var cogType reflect.Type

type TimeAgo struct {
	cog.UXCog
	timeInstance time.Time
}

func NewTimeAgo() *TimeAgo {
	t := &TimeAgo{}
	t.SetCogType(cogType)
	return t
}

func (t *TimeAgo) SetTime(timeInstance time.Time) {
	t.timeInstance = timeInstance
}

func (t *TimeAgo) Start() error {

	if t.timeInstance.IsZero() == true {
		return errors.New("The time instance value has not been set!")
	}

	t.SetProp("timeAgoValue", humanize.Time(t.timeInstance))

	err := t.Render()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	cogType = reflect.TypeOf(TimeAgo{})
}
