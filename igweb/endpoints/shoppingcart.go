package endpoints

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
	"github.com/gorilla/sessions"
)

var SessionStore *sessions.FilesystemStore

func GetShoppingCartItemsEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var cart *models.ShoppingCart
		igwSession, _ := SessionStore.Get(r, "igweb-session")

		if _, ok := igwSession.Values["shoppingCart"]; ok == true {
			// Shopping cart exists in session
			decoder := json.NewDecoder(strings.NewReader(string(igwSession.Values["shoppingCart"].([]byte))))
			err := decoder.Decode(&cart)
			if err != nil {
				log.Print("Encountered error when attempting to decode json data from session: ", err)
			}

			products := env.DB.GetProductsInShoppingCart(cart)
			json.NewEncoder(w).Encode(products)

		} else {
			// Shopping cart doesn't exist in session
			cart = nil
			json.NewEncoder(w).Encode(cart)
			return
		}

	})
}

func AddItemToShoppingCartEndpoint(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		igwSession, _ := SessionStore.Get(r, "igweb-session")
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

		igwSession, _ := SessionStore.Get(r, "igweb-session")
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
		err = json.NewEncoder(b).Encode(cart)
		if err != nil {
			log.Print("Encountered error when attempting to encode cart struct as json data: ", err)
		}
		igwSession.Values["shoppingCart"] = b.Bytes()
		igwSession.Save(r, w)

		w.Write([]byte("OK"))

	})
}

// TODO: move session store initialization to somewhere more appropriate
func init() {
	if _, err := os.Stat("/tmp/igweb-sessions"); os.IsNotExist(err) {
		os.Mkdir("/tmp/igweb-sessions", 711)
	}
	SessionStore = sessions.NewFilesystemStore("/tmp/igweb-sessions", []byte(os.Getenv("IGWEB_HASH_KEY")))

}
