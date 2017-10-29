package datepicker

import (
	"errors"
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

func (d *DatePicker) Start() error {

	if d.Props["datepickerInputID"] == nil {
		return errors.New("Warning: The datePickerInputID prop need to be set!")
	}

	err := d.Render()
	if err != nil {
		return err
	}

	params := &DatePickerParams{Object: js.Global.Get("Object").New()}

	for propName, propValue := range d.Props {
		switch propName {

		case "datepickerInputID":
			inputFieldID := propValue.(string)
			dateInputField := JS.Get("document").Call("getElementById", inputFieldID)
			params.Field = dateInputField

		case "datepickerLabel":
			// Do nothing

		case "datepickerMinDate":
			datepickerMinDate := propValue.(time.Time)
			minDateUnix := datepickerMinDate.Unix()
			params.MinDate = JS.Get("Date").New(minDateUnix * 1000)

		case "datepickerMaxDate":
			datepickerMaxDate := propValue.(time.Time)
			maxDateUnix := datepickerMaxDate.Unix()
			params.MaxDate = JS.Get("Date").New(maxDateUnix * 1000)

		case "datepickerYearRange":
			yearRange := propValue.([]int)
			params.YearRange = yearRange

		default:
			println("Warning: Unknown prop name provided: ", propName)
		}
	}

	d.picker = JS.Get("Pikaday").New(params)

	return nil
}

func init() {
	cogType = reflect.TypeOf(DatePicker{})
}
