package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"go.isomorphicgo.org/go/isokit"
)

func CarsDemoHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env.TemplateSet.Render("carsdemo_page", &isokit.RenderParams{Writer: w, Data: nil})
	})
}
