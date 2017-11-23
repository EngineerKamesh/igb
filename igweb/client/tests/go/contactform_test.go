package main

import (
	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	"github.com/gopherjs/gopherjs/js"
)

type FormParams struct {
	*js.Object
	FirstName   string `js:"firstName"`
	LastName    string `js:"lastName"`
	Email       string `js:"email"`
	MessageBody string `js:"messageBody"`
}

var wait = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Call("waitForSelector", "#primaryContent")
	return nil
})

// Fill out the contact form with a poorly formatted e-mail address
var fillOutContactFormWithPoorlyFormattedEmailAddress = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	params := &FormParams{Object: js.Global.Get("Object").New()}
	params.FirstName = "Isomorphic"
	params.LastName = "Gopher"
	params.Email = "dev@null@test@test.com"
	params.MessageBody = "Sending a contact form submission using CasperJS and PhantomJS"
	this.Call("fill", "#contactForm", params, true)
	return nil
})

// Fill out the contact form partially
var fillOutContactFormPartially = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	params := &FormParams{Object: js.Global.Get("Object").New()}
	params.FirstName = "Isomorphic"
	params.LastName = ""
	params.Email = "devnull@test.com"
	params.MessageBody = ""
	this.Call("fill", "#contactForm", params, true)
	return nil
})

// Fill out the contact form completely
var fillOutContactFormCompletely = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	params := &FormParams{Object: js.Global.Get("Object").New()}
	params.FirstName = "Isomorphic"
	params.LastName = "Gopher"
	params.Email = "devnull@test.com"
	params.MessageBody = "Sending a contact form submission using CasperJS and PhantomJS"
	this.Call("fill", "#contactForm", params, true)
	return nil
})

var casper = js.Global.Get("casper")

func main() {

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Contact Form Test Suite", 4, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/contact", wait)
	})

	// Test validation for poorly formatted e-mail address
	casper.Call("then", fillOutContactFormWithPoorlyFormattedEmailAddress)
	casper.Call("wait", 450, func() {
		casper.Call("capture", "screenshots/contactform_test_invalid_email_error_message.png")
		casper.Get("test").Call("assertSelectorHasText", "#emailError", "The e-mail address entered has an improper syntax", "Display e-mail address syntax error when poorly formatted e-mail entered.")
	})

	// Test validation when contact form is filled out partially
	casper.Call("then", fillOutContactFormPartially)
	casper.Call("wait", 450, func() {
		casper.Call("capture", "screenshots/contactform_test_partially_filled_form_errors.png")
		casper.Get("test").Call("assertSelectorHasText", "#lastNameError", "The last name field is required.", "Display error message when the last name field has not been filled out.")
		casper.Get("test").Call("assertSelectorHasText", "#messageBodyError", "The message area must be filled.", "Display error message when the message body text area has not been filled out.")
	})

	// Test form submission when contact form is filled out completely
	casper.Call("then", fillOutContactFormCompletely)
	casper.Call("wait", 450, func() {
		casper.Call("capture", "screenshots/contactform_confirmation_message.png")
		casper.Get("test").Call("assertSelectorHasText", "#primaryContent h1", "Confirmation", "Display confirmation message after submitting contact form.")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
