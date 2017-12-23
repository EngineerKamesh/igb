package tests

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestContactFormValidation(t *testing.T) {

	testURL := testHost + "/contact"
	expectedTokenMap := map[string]string{"firstName": "The first name field is required.", "/": "The last name field is required.", "email": "The e-mail address entered has an improper syntax.", "messageBody": "The message area must be filled."}

	form := url.Values{}
	form.Add("firstName", "")
	form.Add("lastName", "")
	form.Add("email", "devnull@g@o")
	form.Add("messageBody", "")

	req, err := http.NewRequest("POST", testURL, strings.NewReader(form.Encode()))

	if err != nil {
		t.Errorf("Failed to create new POST request to URL: %s, with error: %s", testURL, err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	hc := http.Client{}
	response, err := hc.Do(req)

	if err != nil {
		t.Errorf("Failed to make POST request to URL: %s, with error: %s", testURL, err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Failed to read response body contents with error: %s", err)
	}

	for k, v := range expectedTokenMap {
		if strings.Contains(string(contents), v) == false {
			t.Errorf("Could not find expected string token: \"%s\" for field \"%s\"", v, k)
		}
	}

}
