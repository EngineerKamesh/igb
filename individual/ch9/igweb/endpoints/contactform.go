package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/forms"
	"github.com/EngineerKamesh/igb/igweb/submissions"
	"github.com/isomorphicgo/isokit"
)

func ContactFormEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var fields map[string]string

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		err = json.Unmarshal(reqBody, &fields)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal json data: ", err)
		}

		formParams := isokit.FormParams{ResponseWriter: w, Request: r, UseFormFieldsForValidation: true, FormFields: fields}
		contactForm := forms.NewContactForm(&formParams)
		validationResult := contactForm.Validate()

		if validationResult == true {
			submissions.ProcessContactForm(env, contactForm)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contactForm.Errors())
	})
}
