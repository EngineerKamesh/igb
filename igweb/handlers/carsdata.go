package handlers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
)

func CarsDataHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var cars []models.Car
		var carsDataBuffer bytes.Buffer

		dec := gob.NewDecoder(&carsDataBuffer)
		body, err := ioutil.ReadAll(r.Body)
		carsDataBuffer = *bytes.NewBuffer(body)
		err = dec.Decode(&cars)

		w.Header().Set("Content-Type", "text/plain")

		if err != nil {
			log.Println(err)
			w.Write([]byte("Something went wrong, look into it"))

		} else {
			fmt.Printf("Cars Data: %#v\n", cars)
			w.Write([]byte("Thanks, I got the slice of cars you sent me!"))
		}

	})
}
