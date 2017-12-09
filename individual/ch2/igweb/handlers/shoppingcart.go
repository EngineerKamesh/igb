package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/isomorphicgo/isokit"
)

func ShoppingCartHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var templateData *templatedata.ShoppingCart
		var cart *models.ShoppingCart
		igwSession, _ := env.Store.Get(r, "igweb-session")

		if _, ok := igwSession.Values["shoppingCart"]; ok == true {
			// Shopping cart exists in session
			decoder := json.NewDecoder(strings.NewReader(string(igwSession.Values["shoppingCart"].([]byte))))
			err := decoder.Decode(&cart)
			if err != nil {
				log.Print("Encountered error when attempting to decode json data from session: ", err)
			}

			products := env.DB.GetProductsInShoppingCart(cart)
			templateData = &templatedata.ShoppingCart{PageTitle: "Shopping Cart", Products: products}

		} else {
			// Shopping cart doesn't exist in session
			templateData = &templatedata.ShoppingCart{PageTitle: "Shopping Cart", Products: nil}
		}

		env.TemplateSet.Render("shopping_cart_page", &isokit.RenderParams{Writer: w, Data: templateData})
	})
}
