package templatefuncs

import (
	"strconv"
	"time"
)

func RubyDate(t time.Time) string {
	layout := time.RubyDate
	return t.Format(layout)
}

func UnixTime(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}
