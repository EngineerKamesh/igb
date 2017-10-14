package timeago

import (
	"reflect"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/uxtoolkit/cog"
)

var cogType reflect.Type

type TimeAgo struct {
	cog.UXCog
}

func NewTimeAgo() *TimeAgo {
	t := &TimeAgo{}
	t.SetCogType(cogType)
	return t
}

func (t *TimeAgo) SetTime(timeInstance time.Time) {
	t.SetProp("timeAgoValue", humanize.Time(timeInstance))
}

func init() {
	cogType = reflect.TypeOf(TimeAgo{})
}
