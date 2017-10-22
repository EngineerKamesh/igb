package templatefuncs

import (
	"os"
	"strconv"
	"time"

	"github.com/isomorphicgo/isokit"
)

func RubyDate(t time.Time) string {
	layout := time.RubyDate
	return t.Format(layout)
}

func UnixTime(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

func IsProduction() bool {
	if isokit.OperatingEnvironment() == isokit.ServerEnvironment {
		return os.Getenv("IGWEB_MODE") == "production"
	} else {
		return false
	}
}
