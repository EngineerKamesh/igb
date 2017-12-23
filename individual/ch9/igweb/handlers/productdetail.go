package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/gorilla/mux"
	"github.com/isomorphicgo/isokit"
)

func ProductDetailHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		productTitle := vars["productTitle"]
		product := env.DB.GetProductDetail(productTitle)
		templateData := &templatedata.ProductDetail{PageTitle: product.Name, Product: product}
		env.TemplateSet.Render("product_detail_page", &isokit.RenderParams{Writer: w, Data: templateData})
	})
}
