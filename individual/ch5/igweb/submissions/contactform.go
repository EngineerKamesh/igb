package submissions

import (
	"log"

	"github.com/EngineerKamesh/igb/igweb/common"

	"github.com/EngineerKamesh/igb/igweb/shared/forms"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
)

func ProcessContactForm(env *common.Env, form *forms.ContactForm) {

	log.Println("Successfully reached process content form function, indicating that the contact form was filled out properly resulting in a positive validation.")

	contactRequest := &models.ContactRequest{FirstName: form.GetFieldValue("firstName"), LastName: form.GetFieldValue("lastName"), Email: form.GetFieldValue("email"), Message: form.GetFieldValue("messageBody")}

	env.DB.CreateContactRequest(contactRequest)
}
