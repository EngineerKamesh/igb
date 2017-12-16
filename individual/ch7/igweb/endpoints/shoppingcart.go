package endpoints

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
)

func GetShoppingCartItemsEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)

		} else {
			// Shopping cart doesn't exist in session
			cart = nil
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cart)
			return
		}

	})
}

func AddItemToShoppingCartEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		igwSession, _ := env.Store.Get(r, "igweb-session")
		decoder := json.NewDecoder(r.Body)
		var m map[string]string
		err := decoder.Decode(&m)
		if err != nil {
			log.Print("Encountered error when attempting to decode json data from request body: ", err)
		}
		defer r.Body.Close()

		var cart *models.ShoppingCart
		if _, ok := igwSession.Values["shoppingCart"]; ok == true {
			// Shopping Cart Exists in Session
			decoder := json.NewDecoder(strings.NewReader(string(igwSession.Values["shoppingCart"].([]byte))))
			err := decoder.Decode(&cart)
			if err != nil {
				log.Print("Encountered error when attempting to decode json data from session: ", err)
			}
		} else {
			// Shopping Cart Doesn't Exist in Session, Create a New One
			cart = models.NewShoppingCart()
		}

		cart.AddItem(m["productSKU"])
		b := new(bytes.Buffer)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(b).Encode(cart)
		if err != nil {
			log.Print("Encountered error when attempting to encode cart struct as json data: ", err)
		}
		igwSession.Values["shoppingCart"] = b.Bytes()
		igwSession.Save(r, w)
		w.Write([]byte("OK"))
	})
}

func RemoveItemFromShoppingCartEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		igwSession, _ := env.Store.Get(r, "igweb-session")
		decoder := json.NewDecoder(r.Body)
		var m map[string]string
		err := decoder.Decode(&m)
		if err != nil {
			log.Print("Encountered error when attempting to decode json data from request body: ", err)
		}
		defer r.Body.Close()

		var cart *models.ShoppingCart
		if _, ok := igwSession.Values["shoppingCart"]; ok == true {
			// Shopping Cart Exists in Session
			decoder := json.NewDecoder(strings.NewReader(string(igwSession.Values["shoppingCart"].([]byte))))
			err := decoder.Decode(&cart)
			if err != nil {
				log.Print("Encountered error when attempting to decode json data from session: ", err)
			}
		} else {
			// Shopping Cart Doesn't Exist in Session, Create a New One
			cart = models.NewShoppingCart()
		}

		for k, v := range cart.Items {
			if v.ProductSKU == m["productSKU"] {
				delete(cart.Items, k)
			}
		}

		b := new(bytes.Buffer)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(b).Encode(cart)
		if err != nil {
			log.Print("Encountered error when attempting to encode cart struct as json data: ", err)
		}
		igwSession.Values["shoppingCart"] = b.Bytes()
		igwSession.Save(r, w)

		w.Write([]byte("OK"))

	})
}
