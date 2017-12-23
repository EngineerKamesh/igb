package forms

import (
	"github.com/EngineerKamesh/igb/igweb/shared/validate"
	"github.com/isomorphicgo/isokit"
)

type ContactForm struct {
	isokit.BasicForm
}

func NewContactForm(formParams *isokit.FormParams) *ContactForm {
	prefillFields := []string{"firstName", "lastName", "email", "messageBody", "byDateInput"}
	fields := make(map[string]string)
	errors := make(map[string]string)
	c := &ContactForm{}
	c.SetPrefillFields(prefillFields)
	c.SetFields(fields)
	c.SetErrors(errors)
	c.SetFormParams(formParams)
	return c
}

func (c *ContactForm) Validate() bool {
	c.RegenerateErrors()
	c.PopulateFields()

	// Check if first name was filled out
	if isokit.FormValue(c.FormParams(), "firstName") == "" {
		c.SetError("firstName", "The first name field is required.")
	}

	// Check if last name was filled out
	if isokit.FormValue(c.FormParams(), "lastName") == "" {
		c.SetError("lastName", "The last name field is required.")
	}

	// Check if message body was filled out
	if isokit.FormValue(c.FormParams(), "messageBody") == "" {
		c.SetError("messageBody", "The message area must be filled.")
	}

	// Check if e-mail address was filled out
	if isokit.FormValue(c.FormParams(), "email") == "" {
		c.SetError("email", "The e-mail address field is required.")
	} else if validate.EmailSyntax(isokit.FormValue(c.FormParams(), "email")) == false {
		// Check e-mail address syntax
		c.SetError("email", "The e-mail address entered has an improper syntax.")

	}

	if len(c.Errors()) > 0 {
		return false

	} else {
		return true
	}
}
