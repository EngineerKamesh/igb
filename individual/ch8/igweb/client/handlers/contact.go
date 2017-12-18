package handlers

import (
	"context"
	"encoding/json"
	"strings"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/forms"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/isomorphicgo/isokit"
)

func ContactHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		contactForm := forms.NewContactForm(nil)
		DisplayContactForm(env, contactForm)
	})
}

func handleContactButtonClickEvent(env *common.Env, event dom.Event, contactForm *forms.ContactForm) {

	event.PreventDefault()
	clientSideValidationResult := contactForm.Validate()

	// Even if we disable the client side validation by forcing the validation result to true, the server-side validation
	// will kick in and keep the contact form validation in check. Go ahead and try it by uncommenting the following line.
	//clientSideValidationResult = true

	if clientSideValidationResult == true {

		contactFormErrorsChannel := make(chan map[string]string)
		go ContactFormSubmissionRequest(contactFormErrorsChannel, contactForm)

		go func() {

			serverContactFormErrors := <-contactFormErrorsChannel
			serverSideValidationResult := len(serverContactFormErrors) == 0

			if serverSideValidationResult == true {
				env.TemplateSet.Render("contact_confirmation_content", &isokit.RenderParams{Data: nil, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})
			} else {
				contactForm.SetErrors(serverContactFormErrors)
				contactForm.DisplayErrors()
			}

		}()

	} else {
		contactForm.DisplayErrors()
	}
}

func ContactFormSubmissionRequest(contactFormErrorsChannel chan map[string]string, contactForm *forms.ContactForm) {

	jsonData, err := json.Marshal(contactForm.Fields())
	if err != nil {
		println("Encountered error: ", err)
		return
	}

	data, err := xhr.Send("POST", "/restapi/contact-form", jsonData)
	if err != nil {
		println("Encountered error: ", err)
		return
	}

	var contactFormErrors map[string]string
	json.NewDecoder(strings.NewReader(string(data))).Decode(&contactFormErrors)

	contactFormErrorsChannel <- contactFormErrors
}

func InitializeContactPage(env *common.Env, contactForm *forms.ContactForm) {

	formElement := env.Document.GetElementByID("contactForm").(*dom.HTMLFormElement)
	contactForm.SetFormParams(&isokit.FormParams{FormElement: formElement})
	contactButton := env.Document.GetElementByID("contactButton").(*dom.HTMLInputElement)
	contactButton.AddEventListener("click", false, func(event dom.Event) {
		handleContactButtonClickEvent(env, event, contactForm)
	})

}

func DisplayContactForm(env *common.Env, contactForm *forms.ContactForm) {
	templateData := &templatedata.Contact{PageTitle: "Contact", Form: contactForm}
	env.TemplateSet.Render("contact_content", &isokit.RenderParams{Data: templateData, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: templateData.PageTitle})
	InitializeContactPage(env, contactForm)
}
