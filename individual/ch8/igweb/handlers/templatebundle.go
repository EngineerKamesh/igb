package handlers

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
)

func TemplateBundleHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var templateContentItemsBuffer bytes.Buffer
		enc := gob.NewEncoder(&templateContentItemsBuffer)
		m := env.TemplateSet.Bundle().Items()
		err := enc.Encode(&m)
		if err != nil {
			log.Print("encoding err: ", err)
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(templateContentItemsBuffer.Bytes())
	})

}
