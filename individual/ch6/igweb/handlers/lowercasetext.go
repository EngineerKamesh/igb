package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/EngineerKamesh/igb/igweb/common"
)

func LowercaseTextTransformHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s string

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		reqBodyString := string(reqBody)

		err = json.Unmarshal([]byte(reqBodyString), &s)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		textBytes, err := json.Marshal(strings.ToLower(s))
		if err != nil {
			log.Print("Encountered error when attempting ot marshal JSON: ", err)
		}
		fmt.Println("textBytes string: ", string(textBytes))
		w.Write(textBytes)

	})

}
