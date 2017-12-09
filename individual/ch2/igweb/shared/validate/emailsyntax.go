package validate

import (
	"log"
	"regexp"
)

const EmailRegex = `(?i)^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,3})+$`

func EmailSyntax(email string) bool {
	validationResult := false
	r, err := regexp.Compile(EmailRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(email)
	return validationResult
}
