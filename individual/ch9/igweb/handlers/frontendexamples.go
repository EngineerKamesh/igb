package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/isomorphicgo/isokit"
)

func FrontEndExamplesHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env.TemplateSet.Render("frontend_examples_page", &isokit.RenderParams{Writer: w, Data: nil})
	})
}
