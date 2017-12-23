package tests

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestContactForm(t *testing.T) {

	testURL := testHost + "/contact"
	expectedTokenString := "The contact form has been successfully completed."

	form := url.Values{}
	form.Add("firstName", "Isomorphic")
	form.Add("lastName", "Gopher")
	form.Add("email", "devnull@test.com")
	form.Add("messageBody", "This is a message sent from the automated contact form test.")

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

	if strings.Contains(string(contents), expectedTokenString) == false {
		t.Errorf("Could not find expected string token: \"%s\"", expectedTokenString)
	}
}
