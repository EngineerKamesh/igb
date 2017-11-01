package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/forms"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/EngineerKamesh/igb/igweb/submissions"
	"github.com/isomorphicgo/isokit"
)

func ContactHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		formParams := isokit.FormParams{ResponseWriter: w, Request: r}
		contactForm := forms.NewContactForm(&formParams)

		switch r.Method {

		case "GET":
			DisplayContactForm(env, contactForm)
		case "POST":
			validationResult := contactForm.Validate()
			if validationResult == true {
				submissions.ProcessContactForm(env, contactForm)
				DisplayConfirmation(env, w, r)
			} else {
				DisplayContactForm(env, contactForm)
			}
		default:
			DisplayContactForm(env, contactForm)
		}

	})
}

func DisplayContactForm(env *common.Env, contactForm *forms.ContactForm) {
	templateData := &templatedata.Contact{PageTitle: "Contact", Form: contactForm}
	env.TemplateSet.Render("contact_page", &isokit.RenderParams{Writer: contactForm.FormParams().ResponseWriter, Data: templateData})
}

func DisplayConfirmation(env *common.Env, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contact-confirmation", 302)
}
