package datepicker

import (
	"reflect"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/uxtoolkit/cog"
)

var cogType reflect.Type
var JS = js.Global

type DatePickerParams struct {
	*js.Object
	Field *js.Object `js:"field"`

	FirstDay  int        `js:"firstDay"`
	MinDate   *js.Object `js:"minDate"`
	MaxDate   *js.Object `js:"maxDate"`
	YearRange []int      `js:"yearRange"`
}

type DatePicker struct {
	cog.UXCog
	picker *js.Object
}

func NewDatePicker() *DatePicker {
	d := &DatePicker{}
	d.SetCogType(cogType)
	return d
}

func (d *DatePicker) Start() {

	if d.Props["datepickerInputID"] == nil {
		return
	}

	inputFieldID := d.Props["datepickerInputID"].(string)
	dateInputField := JS.Get("document").Call("getElementById", inputFieldID)

	params := &DatePickerParams{Object: js.Global.Get("Object").New()}
	params.Field = dateInputField

	if d.Props["datepickerMinDate"] != nil {
		datepickerMinDate := d.Props["datepickerMinDate"].(time.Time)
		minDateUnix := datepickerMinDate.Unix()
		params.MinDate = JS.Get("Date").New(minDateUnix * 1000)
	}

	if d.Props["datepickerMaxDate"] != nil {
		datepickerMaxDate := d.Props["datepickerMaxDate"].(time.Time)
		maxDateUnix := datepickerMaxDate.Unix()
		params.MaxDate = JS.Get("Date").New(maxDateUnix * 1000)
	}

	if d.Props["datepickerYearRange"] != nil {
		yearRange := d.Props["datepickerYearRange"].([]int)
		params.YearRange = yearRange
	}

	d.picker = JS.Get("Pikaday").New(params)
}

func init() {
	cogType = reflect.TypeOf(DatePicker{})
}
