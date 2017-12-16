package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/gorilla/mux"
)

func GetProductDetailEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		productTitle := vars["productTitle"]
		products := env.DB.GetProductDetail(productTitle)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})
}
