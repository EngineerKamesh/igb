package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
)

func GetProductsEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products := env.DB.GetProducts()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})
}
