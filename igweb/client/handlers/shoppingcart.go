package handlers

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/notify"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/isomorphicgo/isokit"
	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

func ShoppingCartHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		/*
			routeVars := ctx.Value(isokit.RouteVarsKey("Vars")).(map[string]string)
			if routeVars != nil {
				println("route vars: ", routeVars)
			}
		*/
		renderShoppingCartItems(env)

	})
}

func renderShoppingCartItems(env *common.Env) {

	productsChannel := make(chan []*models.Product)
	go fetchProductsInShoppingCart(productsChannel)
	products := <-productsChannel
	templateData := &templatedata.ShoppingCart{PageTitle: "Shopping Cart", Products: products}
	env.TemplateSet.Render("shopping_cart_content", &isokit.RenderParams{Data: templateData, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: templateData.PageTitle})
	InitializeShoppingCartPage(env)
	env.Router.RegisterLinks("#primaryContent a")
}

func fetchProductsInShoppingCart(productsChannel chan []*models.Product) {

	data, err := xhr.Send("GET", "/restapi/get-cart-items", nil)
	if err != nil {
		println("Encountered error: ", err)
		println(err)
	}
	var products []*models.Product
	json.NewDecoder(strings.NewReader(string(data))).Decode(&products)

	productsChannel <- products
}

func InitializeShoppingCartPage(env *common.Env) {

	buttons := env.Document.GetElementsByClassName("removeFromCartButton")
	for _, button := range buttons {
		button.AddEventListener("click", false,
			func(event dom.Event) {
				handleRemoveFromCartButtonClickEvent(env, event)

			})
	}

}

func addToCart(productSKU string) {

	m := make(map[string]string)
	m["productSKU"] = productSKU
	jsonData, _ := json.Marshal(m)

	data, err := xhr.Send("PUT", "/restapi/add-item-to-cart", jsonData)
	if err != nil {
		println("Encountered error: ", err)
		println(err)
		notify.Error("Failed to add item to cart!")
		return
	}
	var products []*models.Product
	json.NewDecoder(strings.NewReader(string(data))).Decode(&products)
	notify.Success("Item added to cart")
}

func handleAddToCartButtonClickEvent(event dom.Event) {
	productSKU := event.Target().GetAttribute("data-sku")
	go addToCart(productSKU)
}

func removeFromCart(env *common.Env, productSKU string) {

	m := make(map[string]string)
	m["productSKU"] = productSKU
	jsonData, _ := json.Marshal(m)

	data, err := xhr.Send("DELETE", "/restapi/remove-item-from-cart", jsonData)
	if err != nil {
		println("Encountered error: ", err)
		println(err)
		notify.Error("Failed to remove item from cart!")
		return
	}
	var products []*models.Product
	json.NewDecoder(strings.NewReader(string(data))).Decode(&products)
	renderShoppingCartItems(env)
	notify.Success("Item removed from cart")
}

func handleRemoveFromCartButtonClickEvent(env *common.Env, event dom.Event) {
	productSKU := event.Target().GetAttribute("data-sku")
	go removeFromCart(env, productSKU)
}
