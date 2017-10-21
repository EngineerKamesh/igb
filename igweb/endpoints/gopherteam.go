package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
)

func GetGopherTeamEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gophers := env.DB.GetGopherTeam()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gophers)
	})
}
